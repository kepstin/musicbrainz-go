package mb

type Metadata struct {
	Generator string
	Created string
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
	Id string
	Type string
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
	Id string
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
}

type ReleaseGroup struct {
	Id string
	Type string
	Title *string
	Comment *string
	ArtistCredit *ArtistCredit
	ReleaseList *ReleaseList
	RelationList []*RelationList
	TagList *TagList
	UserTagList *UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Recording struct {
	Id string
	Title *string
	Length uint
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
	Id string
	Type string
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
	Id string "attr"
	Type string "attr"
	Title string
	ArtistCredit *ArtistCredit
	ISWC string
	Disambiguation string
//	AliasList
//	RelationList
//	TagList
//	UserTagList
	Rating *Rating
	UserRating *UserRating
}

type Disc struct {
	Id string "attr"
	Sectors uint
//	ReleaseList
}

type PUID struct {
	Id string "attr"
//	RecordingList
}

type ISRC struct {
	Id string "attr"
//	RecordingList
}

type ArtistCredit struct {
	NameCredit []*NameCredit
}

type Relation struct {
	Type string "attr"
	Target string  "attr"
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
	Type string "attr"
	Script string "attr"
}

type Tag struct {
	Count uint "attr"
	Name string
}

type UserTag struct {
	Name string
}

type Rating struct {
	VotesCount uint
	Rating float32
	RatingPercent float32
}

type UserRating struct {
	Data uint "chardata"
}

type LabelInfo struct {
	CatalogNumber *string
	Label *Label
}

type Medium struct {
	Title *string
	Position uint
	Format *string
	DiscList *DiscList
	TrackList *TrackList
}

type Track struct {
	Position uint
	Title *string
	Length uint
	ArtistCredit *ArtistCredit
	Recording *Recording
	Even bool
}

type Annotation struct {
	Type string
	Entity string
	Name string
	Text string
}

type CDStub struct {
	Id string "attr"
	Title string
	Artist string
	Barcode string
	Comment string
//	NonMBTrackList *NonMBTrackList
}

type FreeDBDisc struct {
	Id string "attr"
	Title string
	Artist string
	Category string
	Year string
//	NonMBTrackList *NonMBTrackList
}

type NonMBTrack struct {
	Title string
	Artist string
	Length uint
}

type Collection struct {
	Id string "attr"
	Name string
	Editor string
//	ReleaseList *ReleaseList
}

type ArtistList struct {
	Artist []Artist
}

type MediumList struct {
	List
	TrackCount uint
	Medium []*Medium
}

type ReleaseList struct {
	List
	Release []Release
}

type ReleaseGroupList struct {
	List
	ReleaseGroup []ReleaseGroup
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

type NameCredit struct {
	JoinPhrase string
	Name *string
	Artist *Artist
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
	Count uint
	Offset uint
}

type LifeSpan struct {
	Begin *string
	End *string
}
