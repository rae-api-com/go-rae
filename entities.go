package rae

//go:generate easyjson

type WordCategory string

const (
	CategoryArticle      WordCategory = "article" // articulo
	CategoryNoun         WordCategory = "noun"    // sustantivo
	CategoryPronoun      WordCategory = "pronoun"
	CategoryAdjective    WordCategory = "adjective"
	CategoryVerb         WordCategory = "verb"
	CategoryAdverb       WordCategory = "adverb"
	CategoryPreposition  WordCategory = "preposition"
	CategoryConjunction  WordCategory = "conjunction"
	CategoryInterjection WordCategory = "interjection"
)

type VerbCategory string

const (
	VerbCategoryTransitive   VerbCategory = "transitive"   // Verbos transitivos
	VerbCategoryIntransitive VerbCategory = "intransitive" // Verbos intransitivos
	VerbCategoryCopulative   VerbCategory = "copulative"   // Verbos copulativos
	VerbCategoryReflexive    VerbCategory = "reflexive"    // Verbos reflexivos
	VerbCategoryDefective    VerbCategory = "defective"    // Verbos defectivos
	VerbCategoryPronominal   VerbCategory = "pronominal"   // Verbos recíprocos
	VerbCategoryAuxiliary    VerbCategory = "auxiliary"    // Verbos auxiliares
	VerbCategoryPredicative  VerbCategory = "predicative"  // Verbos predicativos
)

type ArticleCategory string

const (
	ArticleCategoryDefinite   ArticleCategory = "definite"
	ArticleCategoryIndefinite ArticleCategory = "indefinite"
	ArticleCategoryNeuter     ArticleCategory = "neuter"
)

// Usage represents the usage type.
type Usage string

const (
	UsageCommon     Usage = "common"
	UsageRare       Usage = "rare"
	UsageOutdated   Usage = "outdated"
	UsageColloquial Usage = "colloquial"
	UsageObsolete   Usage = "obsolete" // desuso
	UsageUnknown    Usage = "unknown"
)

type OriginType string

const (
	OriginLatin     OriginType = "lat"
	OriginUncertain OriginType = "uncertain"
)

type VoiceType string

const (
	VoiceOnomatopoeic VoiceType = "onomatopoeic"
	VoiceExpressive   VoiceType = "expressive"
)

type Gender string

const (
	GenderMasculine Gender = "masculine"
	GenderFeminine  Gender = "feminine"
	GenderBoth      Gender = "masculine_and_feminine"
	GenderUnknown   Gender = "unknown"
)

type VerbalMode string

const (
	VerbalModeIndicative  VerbalMode = "indicative"
	VerbalModeSubjunctive VerbalMode = "subjunctive"
	VerbalModeImperative  VerbalMode = "imperative"
	VerbalModeNonPersonal VerbalMode = "nonpersonal"
)

// Definition represents a word definition.
//
//easyjson:json
type Definition struct {
	Raw           string        `json:"raw"`
	MeaningNumber int           `json:"meaning_number"`
	Category      WordCategory  `json:"category"`
	VerbCategory  *VerbCategory `json:"verb_category,omitempty"`
	Gender        *Gender       `json:"gender,omitempty"`
	Article       *Article      `json:"article,omitempty"`
	Usage         Usage         `json:"usage"`
	Description   string        `json:"description"`
	Synonyms      []string      `json:"synonyms"`
	Antonyms      []string      `json:"antonyms"`
}

type Origin struct {
	Raw   string     `json:"raw"`
	Type  OriginType `json:"type"`
	Voice VoiceType  `json:"voice"`
	Text  string     `json:"text"`
}

// WordEntry represents an entry for a word.

type (
	//easyjson:json
	ConjugationNonPersonal struct {
		Infinitive         string `json:"infinitive"`
		Participle         string `json:"participle"`
		Gerund             string `json:"gerund"`
		CompoundInfinitive string `json:"compound_infinitive"`
		CompoundGerund     string `json:"compound_gerund"`
	}

	//easyjson:json
	ConjugationIndicative struct {
		Present            Conjugation `json:"present"`
		PresentPerfect     Conjugation `json:"present_perfect"`     // Pretérito perfecto compuesto / Antepresente
		Imperfect          Conjugation `json:"imperfect"`           // Pretérito imperfecto / Copretérito
		PastPerfect        Conjugation `json:"past_perfect"`        // Pretérito pluscuamperfecto / Antecopretérito
		Preterite          Conjugation `json:"preterite"`           // Pretérito perfecto simple / Pretérito
		PastAnterior       Conjugation `json:"past_anterior"`       // Pretérito anterior / Antepretérito
		Future             Conjugation `json:"future"`              // Futuro simple
		FuturePerfect      Conjugation `json:"future_perfect"`      // Futuro compuesto / Antefuturo
		Conditional        Conjugation `json:"conditional"`         // Condicional simple / Pospretérito
		ConditionalPerfect Conjugation `json:"conditional_perfect"` // Condicional compuesto / Antepospretérito
	}

	//easyjson:json
	ConjugationSubjunctive struct {
		Present        Conjugation `json:"present"`
		PresentPerfect Conjugation `json:"present_perfect"` // Pretérito perfecto compuesto / Antepresente
		Imperfect      Conjugation `json:"imperfect"`       // Pretérito imperfecto / Pretérito
		PastPerfect    Conjugation `json:"past_perfect"`    // Pretérito pluscuamperfecto / Antepretérito
		Future         Conjugation `json:"future"`          // Futuro simple
		FuturePerfect  Conjugation `json:"future_perfect"`  // Futuro compuesto / Antefuturo
	}

	//easyjson:json
	ConjugationImperative struct {
		SingularSecondPerson       string `json:"singular_second_person"`
		SingularFormalSecondPerson string `json:"singular_formal_second_person"`
		PluralSecondPerson         string `json:"plural_second_person"`
		PluralFormalSecondPerson   string `json:"plural_formal_second_person"`
	}
)

//easyjson:json
type Conjugations struct {
	ConjugationNonPersonal ConjugationNonPersonal `json:"non_personal"`
	ConjugationIndicative  ConjugationIndicative  `json:"indicative"`
	ConjugationSubjunctive ConjugationSubjunctive `json:"subjunctive"`
	ConjugationImperative  ConjugationImperative  `json:"imperative"`
}

//easyjson:json
type WordEntry struct {
	Word        string    `json:"word"`
	Meanings    []Meaning `json:"meanings"`
	Suggestions []string  `json:"suggestions"`
}

//easyjson:json
type AdditionalSense struct {
	Number     int        `json:"number"`
	Definition string     `json:"definition"`
	Locutions  []Locution `json:"locutions"`
}

//easyjson:json
type Locution struct {
	Text       string `json:"text"`
	Definition string `json:"definition"`
}

//easyjson:json
type Conjugation struct {
	SingularFirstPerson        string `json:"singular_first_person"`
	SingularSecondPerson       string `json:"singular_second_person"`
	SingularFormalSecondPerson string `json:"singular_formal_second_person"`
	SingularThirdPerson        string `json:"singular_third_person"`
	PluralFirstPerson          string `json:"plural_first_person"`
	PluralSecondPerson         string `json:"plural_second_person"`
	PluralFormalSecondPerson   string `json:"plural_formal_second_person"`
	PluralThirdPerson          string `json:"plural_third_person"`
}

//easyjson:json
type Article struct {
	Category ArticleCategory `json:"category"`
	Gender   Gender          `json:"gender"`
}

//easyjson:json
type Meaning struct {
	Origin       *Origin       `json:"origin,omitempty"`
	Definitions  []Definition  `json:"senses"`
	Conjugations *Conjugations `json:"conjugations,omitempty"`
}
