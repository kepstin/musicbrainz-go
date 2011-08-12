package mb

import (
	"io"
	"os"
	"strconv"
	"xml"
)

func Parse(reader io.Reader) (*Metadata, os.Error) {
	parser := xml.NewParser(reader)
	var metadata *Metadata
	var err os.Error
	for tok, err := parser.Token(); err == nil; tok, err = parser.Token() {
		switch xmltok := tok.(type) {
		case xml.StartElement:
			switch xmltok.Name.Local {
			case "metadata":
				metadata, err = parseMetadata(parser, xmltok)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil && err != os.EOF {
		return nil, err
	}
	if metadata == nil {
		return nil, os.NewError("Missing <metadata> tag")
	}
	return metadata, nil
}

func parseMetadata(parser *xml.Parser, se xml.StartElement) (*Metadata, os.Error) {
	metadata := new(Metadata)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "generator":
			metadata.Generator = attr.Value
		case "created":
			metadata.Created = attr.Value
		}
	}
	for tok, err := parser.Token(); err == nil; tok, err = parser.Token() {
		switch xmltok := tok.(type) {
		case xml.EndElement:
			switch xmltok.Name.Local {
			case "metadata":
				return metadata, nil
			}
		case xml.StartElement:
			switch xmltok.Name.Local {
			case "artist":
				metadata.Artist, err = parseArtist(parser, xmltok)
			case "release":
				metadata.Release, err = parseRelease(parser, xmltok)
			case "release-group":
				metadata.ReleaseGroup, err = parseReleaseGroup(parser, xmltok)
			case "recording":
				metadata.Recording, err = parseRecording(parser, xmltok)
			case "label":
				metadata.Label, err = parseLabel(parser, xmltok)
			case "work":
				metadata.Work, err = parseWork(parser, xmltok)
			case "puid":
				metadata.PUID, err = parsePUID(parser, xmltok)
			case "isrc":
				metadata.ISRC, err = parseISRC(parser, xmltok)
			case "disc":
				metadata.Disc, err = parseDisc(parser, xmltok)
			case "rating":
				metadata.Rating, err = parseRating(parser, xmltok)
			case "user-rating":
				metadata.UserRating, err = parseUserRating(parser, xmltok)
			case "collection":
				metadata.Collection, err = parseCollection(parser, xmltok)
			case "artist-list":
				metadata.ArtistList, err = parseArtistList(parser, xmltok)
			case "release-list":
				metadata.ReleaseList, err = parseReleaseList(parser, xmltok)
			case "release-group-list":
				metadata.ReleaseGroupList, err = parseReleaseGroupList(parser, xmltok)
			case "recording-list":
				metadata.RecordingList, err = parseRecordingList(parser)
			case "label-list":
				metadata.LabelList, err = parseLabelList(parser)
			case "work-list":
				metadata.WorkList, err = parseWorkList(parser)
			case "isrc-list":
				metadata.ISRCList, err = parseISRCList(parser)
			case "annotation-list":
				metadata.AnnotationList, err = parseAnnotationList(parser)
			case "cdstub-list":
				metadata.CDStubList, err = parseCDStubList(parser)
			case "freedb-disc-list":
				metadata.FreeDBDiscList, err = parseFreeDBDiscList(parser)
			case "tag-list":
				metadata.TagList, err = parseTagList(parser)
			case "user-tag-list":
				metadata.UserTagList, err = parseUserTagList(parser)
			case "collection-list":
				metadata.CollectionList, err = parseCollectionList(parser)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </metadata> tag")
}

func parseArtist(p *xml.Parser, t xml.StartElement) (*Artist, os.Error) {
	a := new(Artist)
	for _, attr := range t.Attr {
		switch attr.Name.Local {
		case "id":
			a.Id = attr.Value
		case "type":
			a.Type = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "artist":
				return a, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "name":
				a.Name, err = parseXmlString(p, tt)
			case "sort-name":
				a.SortName, err = parseXmlString(p, tt)
			case "gender":
				a.Gender, err = parseXmlString(p, tt)
			case "country":
				a.Country, err = parseXmlString(p, tt)
			case "disambiguation":
				a.Disambiguation, err = parseXmlString(p, tt)
			case "life-span":
				a.LifeSpan, err = parseLifeSpan(p)
			case "alias-list":
				a.AliasList, err = parseAliasList(p)
			case "recording-list":
				a.RecordingList, err = parseRecordingList(p)
			case "release-list":
				a.ReleaseList, err = parseReleaseList(p, tt)
			case "release-group-list":
				a.ReleaseGroupList, err = parseReleaseGroupList(p, tt)
			case "label-list":
				a.LabelList, err = parseLabelList(p)
			case "work-list":
				a.WorkList, err = parseWorkList(p)
			case "relation-list":
				rl, err := parseRelationList(p)
				if err == nil {
					a.RelationList = append(a.RelationList, rl)
				}
			case "tag-list":
				a.TagList, err = parseTagList(p)
			case "user-tag-list":
				a.UserTagList, err = parseUserTagList(p)
			case "rating":
				a.Rating, err = parseRating(p, tt)
			case "user-rating":
				a.UserRating, err = parseUserRating(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </artist> tag")
}

func parseRelease(p *xml.Parser, se xml.StartElement) (*Release, os.Error) {
	r := new(Release)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "id":
			r.Id = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "release":
				return r, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "title":
				r.Title, err = parseXmlString(p, tt)
			case "status":
				r.Status, err = parseXmlString(p, tt)
			case "quality":
				r.Quality, err = parseXmlString(p, tt)
			case "disambiguation":
				r.Disambiguation, err = parseXmlString(p, tt)
			case "packaging":
				r.Packaging, err = parseXmlString(p, tt)
			case "text-representation":
				r.TextRepresentation, err = parseTextRepresentation(p)
			case "artist-credit":
				r.ArtistCredit, err = parseArtistCredit(p)
			case "release-group":
				r.ReleaseGroup, err = parseReleaseGroup(p, tt)
			case "date":
				r.Date, err = parseXmlString(p, tt)
			case "country":
				r.Country, err = parseXmlString(p, tt)
			case "barcode":
				r.Barcode, err = parseXmlString(p, tt)
			case "asin":
				r.ASIN, err = parseXmlString(p, tt)
			case "label-info-list":
				r.LabelInfoList, err = parseLabelInfoList(p, tt)
			case "medium-list":
				r.MediumList, err = parseMediumList(p, tt)
			case "relation-list":
				rl, err := parseRelationList(p)
				if err == nil {
					r.RelationList = append(r.RelationList, rl)
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </release> tag")
}

func parseReleaseGroup(p *xml.Parser, se xml.StartElement) (*ReleaseGroup, os.Error) {
	rg := new(ReleaseGroup)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "id":
			rg.Id = attr.Value
		case "type":
			rg.Type = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "release-group":
				return rg, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "title":
				rg.Title, err = parseXmlString(p, tt)
			case "comment":
				rg.Comment, err = parseXmlString(p, tt)
			case "artist-credit":
				rg.ArtistCredit, err = parseArtistCredit(p)
			case "release-list":
				rg.ReleaseList, err = parseReleaseList(p, tt)
			case "relation-list":
				rl, err := parseRelationList(p)
				if err == nil {
					rg.RelationList = append(rg.RelationList, rl)
				}
			case "tag-list":
				rg.TagList, err = parseTagList(p)
			case "user-tag-list":
				rg.UserTagList, err = parseUserTagList(p)
			case "rating":
				rg.Rating, err = parseRating(p, tt)
			case "user-rating":
				rg.UserRating, err = parseUserRating(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </release-group> tag")
}

func parseRecording(p *xml.Parser, se xml.StartElement) (*Recording, os.Error) {
	r := new(Recording)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "id":
			r.Id = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "recording":
				return r, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "title":
				r.Title, err = parseXmlString(p, tt)
			case "length":
				r.Length, err = parseXmlUint(p, tt)
			case "disambiguation":
				r.Disambiguation, err = parseXmlString(p, tt)
			case "artist-credit":
				r.ArtistCredit, err = parseArtistCredit(p)
			case "release-list":
				r.ReleaseList, err = parseReleaseList(p, tt)
			case "puid-list":
				r.PUIDList, err = parsePUIDList(p, tt)
			case "isrc-list":
				r.ISRCList, err = parseISRCList(p)
			case "relation-list":
				rl, err := parseRelationList(p)
				if err == nil {
					r.RelationList = append(r.RelationList, rl)
				}
			case "tag-list":
				r.TagList, err = parseTagList(p)
			case "user-tag-list":
				r.UserTagList, err = parseUserTagList(p)
			case "rating":
				r.Rating, err = parseRating(p, tt)
			case "user-rating":
				r.UserRating, err = parseUserRating(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </recording> tag")
}

func parseLabel(p *xml.Parser, se xml.StartElement) (*Label, os.Error) {
	l := new(Label)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "id":
			l.Id = attr.Value
		case "type":
			l.Type = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "label":
				return l, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "name":
				l.Name, err = parseXmlString(p, tt)
			case "sort-name":
				l.SortName, err = parseXmlString(p, tt)
			case "label-code":
				l.LabelCode, err = parseXmlUint(p, tt)
			case "disambiguation":
				l.Disambiguation, err = parseXmlString(p, tt)
			case "country":
				l.Country, err = parseXmlString(p, tt)
			case "life-span":
				l.LifeSpan, err = parseLifeSpan(p)
			case "alias-list":
				l.AliasList, err = parseAliasList(p)
			case "release-list":
				l.ReleaseList, err = parseReleaseList(p, tt)
			case "relation-list":
				rl, err := parseRelationList(p)
				if err == nil {
					l.RelationList = append(l.RelationList, rl)
				}
			case "tag-list":
				l.TagList, err = parseTagList(p)
			case "user-tag-list":
				l.UserTagList, err = parseUserTagList(p)
			case "rating":
				l.Rating, err = parseRating(p, tt)
			case "user-rating":
				l.UserRating, err = parseUserRating(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </release-group> tag")
}

func parseWork(parser *xml.Parser, se xml.StartElement) (*Work, os.Error) {
	parser.Skip()
	return nil, nil
}

func parsePUID(parser *xml.Parser, se xml.StartElement) (*PUID, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseISRC(parser *xml.Parser, se xml.StartElement) (*ISRC, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseArtistCredit(p *xml.Parser) (*ArtistCredit, os.Error) {
	ac := new(ArtistCredit)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "artist-credit":
				return ac, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "name-credit":
				nc, err := parseNameCredit(p, tt)
				if err == nil {
					ac.NameCredit = append(ac.NameCredit, nc)
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </artist-credit> tag")
}

func parseDisc(parser *xml.Parser, se xml.StartElement) (*Disc, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseRating(p *xml.Parser, se xml.StartElement) (*Rating, os.Error) {
	r := new(Rating)
	var err os.Error
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "votes-count":
			r.VotesCount, err = strconv.Atoui(attr.Value)
		}
		if err != nil {
			return nil, err
		}
	}
	s, err := parseXmlString(p, se)
	if err != nil {
		return nil, err
	}
	r.Rating, err = strconv.Atof32(*s)
	if err != nil {
		return nil, err
	}
	r.RatingPercent = r.Rating * 20
	return r, nil
}

func parseUserRating(parser *xml.Parser, se xml.StartElement) (*UserRating, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseLabelInfo(p *xml.Parser) (*LabelInfo, os.Error) {
	li := new(LabelInfo)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "label-info":
				return li, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "catalog-number":
				li.CatalogNumber, err = parseXmlString(p, tt)
			case "label":
				li.Label, err = parseLabel(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </label-info> tag")
}

func parseMedium(p *xml.Parser) (*Medium, os.Error) {
	m := new(Medium)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "medium":
				return m, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "title":
				m.Title, err = parseXmlString(p, tt)
			case "position":
				m.Position, err = parseXmlUint(p, tt)
			case "format":
				m.Format, err = parseXmlString(p, tt)
			case "disc-list":
				m.DiscList, err = parseDiscList(p)
			case "track-list":
				m.TrackList, err = parseTrackList(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </medium> tag")
}

func parseTrack(p *xml.Parser) (*Track, os.Error) {
	tr := new(Track)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "track":
				return tr, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "position":
				tr.Position, err = parseXmlUint(p, tt)
				if tr.Position%2 == 0 {
					tr.Even = true
				}
			case "title":
				tr.Title, err = parseXmlString(p, tt)
			case "length":
				tr.Length, err = parseXmlUint(p, tt)
			case "artist-credit":
				tr.ArtistCredit, err = parseArtistCredit(p)
			case "recording":
				tr.Recording, err = parseRecording(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </track> tag")
}

func parseCollection(parser *xml.Parser, se xml.StartElement) (*Collection, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseArtistList(parser *xml.Parser, se xml.StartElement) (*ArtistList, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseMediumList(p *xml.Parser, se xml.StartElement) (*MediumList, os.Error) {
	ml := new(MediumList)
	err := ml.parseListAttributes(p, se)
	if err != nil {
		return nil, err
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "medium-list":
				return ml, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "track-count":
				ml.TrackCount, err = parseXmlUint(p, tt)
			case "medium":
				m, err := parseMedium(p)
				if err == nil {
					ml.Medium = append(ml.Medium, m)
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </medium-list> tag")
}

func parseReleaseList(parser *xml.Parser, se xml.StartElement) (*ReleaseList, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseReleaseGroupList(parser *xml.Parser, se xml.StartElement) (*ReleaseGroupList, os.Error) {
	parser.Skip()
	return nil, nil
}

func parseAliasList(p *xml.Parser) (*AliasList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseRecordingList(p *xml.Parser) (*RecordingList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseTrackList(p *xml.Parser, se xml.StartElement) (*TrackList, os.Error) {
	tl := new(TrackList)
	err := tl.parseListAttributes(p, se)
	if err != nil {
		return nil, err
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "track-list":
				return tl, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "track":
				tr, err := parseTrack(p)
				if err == nil {
					tl.Track = append(tl.Track, tr)
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </track-list> tag")
}

func parseLabelList(p *xml.Parser) (*LabelList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseLabelInfoList(p *xml.Parser, se xml.StartElement) (*LabelInfoList, os.Error) {
	lil := new(LabelInfoList)
	err := lil.parseListAttributes(p, se)
	if err != nil {
		return nil, err
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "label-info-list":
				return lil, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "label-info":
				li, err := parseLabelInfo(p)
				if err == nil {
					lil.LabelInfo = append(lil.LabelInfo, li)
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </label-info-list> tag")
}

func parseWorkList(p *xml.Parser) (*WorkList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseISRCList(p *xml.Parser) (*ISRCList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseRelationList(p *xml.Parser) (*RelationList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseAnnotationList(p *xml.Parser) (*AnnotationList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseCDStubList(p *xml.Parser) (*CDStubList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseFreeDBDiscList(p *xml.Parser) (*FreeDBDiscList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseDiscList(p *xml.Parser) (*DiscList, os.Error) {
	p.Skip()
	return nil, nil
}

func parsePUIDList(p *xml.Parser, se xml.StartElement) (*PUIDList, os.Error) {
	pl := new(PUIDList)
	err := pl.parseListAttributes(p, se)
	if err != nil {
		return nil, err
	}
	return pl, nil
}

func parseTagList(p *xml.Parser) (*TagList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseUserTagList(p *xml.Parser) (*UserTagList, os.Error) {
	p.Skip()
	return nil, nil
}

func parseCollectionList(p *xml.Parser) (*CollectionList, os.Error) {
	p.Skip()
	return nil, nil
}

func (l *List) parseListAttributes(p *xml.Parser, se xml.StartElement) os.Error {
	return nil
}

func parseLifeSpan(p *xml.Parser) (*LifeSpan, os.Error) {
	ls := new(LifeSpan)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "life-span":
				return ls, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "begin":
				ls.Begin, err = parseXmlString(p, tt)
			case "end":
				ls.End, err = parseXmlString(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </life-span> tag")
}

func parseTextRepresentation(p *xml.Parser) (*TextRepresentation, os.Error) {
	tr := new(TextRepresentation)
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "text-representation":
				return tr, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "language":
				tr.Language, err = parseXmlString(p, tt)
			case "script":
				tr.Script, err = parseXmlString(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </text-representation> tag")
}

func parseNameCredit(p *xml.Parser, se xml.StartElement) (*NameCredit, os.Error) {
	nc := new(NameCredit)
	for _, attr := range se.Attr {
		switch attr.Name.Local {
		case "joinphrase":
			nc.JoinPhrase = attr.Value
		}
	}
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.EndElement:
			switch tt.Name.Local {
			case "name-credit":
				return nc, nil
			}
		case xml.StartElement:
			switch tt.Name.Local {
			case "name":
				nc.Name, err = parseXmlString(p, tt)
			case "artist":
				nc.Artist, err = parseArtist(p, tt)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, os.NewError("Missing </name-credit> tag")
}

func parseXmlString(p *xml.Parser, se xml.StartElement) (*string, os.Error) {
	ret := ""
	tag := se.Name.Local
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch tt := t.(type) {
		case xml.CharData:
			ret = ret + (string)(([]byte)(tt))
		case xml.EndElement:
			if tt.Name.Local == tag {
				return &ret, nil
			} else {
				return nil, os.NewError("Unexpected tag </" + tt.Name.Local + ">")
			}
		}
	}
	return nil, os.NewError("Missing end tag </" + tag + ">")
}

func parseXmlUint(p *xml.Parser, se xml.StartElement) (uint, os.Error) {
	s, err := parseXmlString(p, se)
	if err != nil {
		return 0, err
	}
	return strconv.Atoui(*s)
}
