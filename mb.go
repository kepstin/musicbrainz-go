package mb

import "xml"

type Metadata struct {
	XMLName xml.Name `metadata`
	Generator string `attr`
	Created string `attr`
	Artist *Artist
	Release *Release
	ReleaseGroup *ReleaseGroup
	Recording *Recording
	Label *Label
	Work *Work
	PUID *PUID
	ISRC *ISRC
	Disc *Disc
	Rating *Rating
	UserRating *UserRating
	Collection *Collection
	ArtistList *ArtistList
	ReleaseList *ReleaseList
	ReleaseGroupList *ReleaseGroupList
	RecordingList *RecordingList
	LabelList *LabelList
	WorkList *WorkList
	ISRCList *ISRCList
	AnnotationList *AnnotationList
	CDStubList *CDStubList
	FreeDBDiscList *FreeDBDiscList
	TagList *TagList
	UserTagList *UserTagList
	CollectionList *CollectionList
}

type Artist struct {
	XMLName xml.Name `artist`
	Id string `attr`
	Type string `attr`
	Name *string
	SortName *string
	Gender *string
	Country *string
	Disambiguation *string
	LifeSpan *LifeSpan
	AliasList *AliasList
	RecordingList *RecordingList
	ReleaseList *ReleaseList
	ReleaseGroupList *ReleaseGroupList
	LabelList *LabelList
	WorkList *WorkList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Release struct {
	XMLName xml.Name `release`
	Id string `attr`
	Title *string
	Status *string
	Quality *string
	Disambiguation *string
	Packaging *string
	TextRepresentation *TextRepresentation
	ArtistCredit *ArtistCredit
	ReleaseGroup *ReleaseGroup
	Date *string
	Country *string
	Barcode *string
	ASIN *string
	LabelInfoList *LabelInfoList
	MediumList *MediumList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
}

type ReleaseGroup struct {
	XMLName xml.Name `release-group`
	Id string `attr`
	Type string `attr`
	Title *string
	Disambiguation *string
	Comment *string
	FirstReleaseDate *string
	ArtistCredit *ArtistCredit
	ReleaseList *ReleaseList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Recording struct {
	XMLName xml.Name `recording`
	Id string `attr`
	Title *string
	Length uint // milliseconds
	Disambiguation *string
	ArtistCredit *ArtistCredit
	ReleaseList *ReleaseList
	PUIDList *PUIDList
	ISRCList *ISRCList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Label struct {
	XMLName xml.Name `label`
	Id string `attr`
	Type string `attr`
	Name *string
	SortName *string
	LabelCode uint
	Disambiguation *string
	Country *string
	LifeSpan *LifeSpan
	AliasList *AliasList
	ReleaseList *ReleaseList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Work struct {
	XMLName xml.Name `work`
	Id string `attr`
	Type string `attr`
	Title *string
	ArtistCredit *ArtistCredit // Not used
	ISWC *string
	Disambiguation *string
	AliasList *AliasList
	RelationList *RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Disc struct {
	XMLName xml.Name `disc`
	Id string `attr`
	Sectors uint
	ReleaseList *ReleaseList
}

type PUID struct {
	XMLName xml.Name `puid`
	Id string `attr`
	RecordingList *RecordingList
}

type ISRC struct {
	XMLName xml.Name `isrc`
	Id string `attr`
	RecordingList *RecordingList
}

type ArtistCredit struct {
	XMLName xml.Name `artist-credit`
	NameCredit []*NameCredit
}
type NameCredit struct {
	XMLName xml.Name `name-credit`
	JoinPhrase string `attr`
	Name *string
	Artist *Artist
}


type Relation struct {
	XMLName xml.Name `relation`
	Type string `attr`
	Target string  `attr`
	Direction string
	Begin string
	End string
	Artist *Artist
	Release *Release
	ReleaseGroup *ReleaseGroup
	Recording *Recording
	Label *Label
	Work *Work
}

type Alias struct {
	XMLName xml.Name `alias`
	Locale string `attr`
	Text string `chardata`
}

type Tag struct {
	XMLName xml.Name `tag`
	Count uint `attr`
	Name string
}

type UserTag struct {
	XMLName xml.Name `user-tag`
	Name string
}

type Rating struct {
	XMLName xml.Name `rating`
	VotesCount uint
	Rating float32 `chardata`
}

type UserRating struct {
	XMLName xml.Name `user-rating`
	Rating uint `chardata`
}

type LabelInfo struct {
	XMLName xml.Name `label-info`
	CatalogNumber *string
	Label *Label
}

type Medium struct {
	XMLName xml.Name `medium`
	Title *string
	Position uint
	Format *string
	DiscList *DiscList
	TrackList *TrackList
}

type Track struct {
	XMLName xml.Name `track`
	Position uint
	Title *string
	Length uint
	ArtistCredit *ArtistCredit
	Recording *Recording
}

type Annotation struct {
	XMLName xml.Name `annotation`
	Type string `attr`
	Entity string
	Name string
	Text string
}

type CDStub struct {
	XMLName xml.Name `cdstub`
	Id string `attr`
	Title string
	Artist *string
	Barcode *string
	Comment *string
	NonMBTrackList *NonMBTrackList
}

type FreeDBDisc struct {
	XMLName xml.Name `freedb-disc`
	Id string `attr`
	Title string
	Artist *string
	Category *string
	Year *string
	NonMBTrackList *NonMBTrackList
}

type NonMBTrack struct {
	XMLName xml.Name `track`
	Title string
	Artist *string
	Length uint
}

type Collection struct {
	XMLName xml.Name `collection`
	Id string `attr`
	Name string
	Editor *string
	ReleaseList *ReleaseList
}

type ArtistList struct {
	XMLName xml.Name `artist-list`
	Artist []*Artist
}

type MediumList struct {
	XMLName xml.Name `medium-list`
	List
	TrackCount uint
	Medium []*Medium
}

type ReleaseList struct {
	XMLName xml.Name `release-list`
	List
	Release []*Release
}

type ReleaseGroupList struct {
	XMLName xml.Name `release-group-list`
	List
	ReleaseGroup []*ReleaseGroup
}

type AliasList struct {
	List
	Alias []Alias
}

type RecordingList struct {
	List
	Recording []Recording
}

type LabelList struct {
	List
	Label []Label
}

type LabelInfoList struct {
	List
	LabelInfo []*LabelInfo
}

type WorkList struct {
	List
	Work []Work
}

type AnnotationList struct {
	List
	Annotation []Annotation
}

type CDStubList struct {
	List
	CDStub []CDStub
}

type NonMBTrackList struct {
	XMLName xml.Name `track-list`
	List
	Track []*NonMBTrack
}

type FreeDBDiscList struct {
	List
	FreeDBDisc []FreeDBDisc
}

type DiscList struct {
	List
	Disc []*Disc
}

type PUIDList struct {
	List
	PUID []*PUID
}

type ISRCList struct {
	List
	ISRC []ISRC
}

type RelationList struct {
	TargetType string
	List
	Relation []Relation
}

type TagList struct {
	List
	Tag []Tag
}

type UserTagList struct {
	List
	UserTag []UserTag
}

type CollectionList struct {
	List
	Collection []Collection
}

type TextRepresentation struct {
	Language *string
	Script *string
}

type TrackList struct {
	List
	Track []*Track
}

type List struct {
	Count uint `attr`
	Offset uint `attr`
}

type LifeSpan struct {
	Begin *string
	End *string
}
