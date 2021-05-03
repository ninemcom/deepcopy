package deepcopy

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCopy(t *testing.T) {
	p := NewComplicatedMan()
	c := ComplicatedMan{}

	Clone(&c, &p)

	t.Log(p)
	t.Log(c)

	pjb, err := json.Marshal(&p)
	assert.Nil(t, err)
	cjb, err := json.Marshal(&c)
	assert.Nil(t, err)

	pjbStr, cjbStr := string(pjb), string(cjb)
	assert.Equal(t, pjbStr, cjbStr)

	pRawStr, cRawStr := fmt.Sprintf("%v", p), fmt.Sprintf("%v", c)
	splitedP, splitedC := strings.Split(pRawStr, " "), strings.Split(cRawStr, " ")
	assert.Equal(t, len(splitedP), len(splitedC))

	numFoundPtr := 0
	for i := range splitedP {
		if strings.HasPrefix(splitedP[i], "0x") || strings.HasPrefix(splitedP[i], "[0x") {
			// point must not be same
			assert.NotEqual(t, splitedP[i], splitedC[i])
			numFoundPtr++
		} else {
			assert.Equal(t, splitedP[i], splitedC[i])
		}
	}
	t.Log("NumPtr:", numFoundPtr)
}

type ComplicatedMan struct {
	Name     string
	FakeName *string

	Age     int
	FakeAge *int

	Birth     time.Time
	FakeBirth *time.Time

	Schedule     []time.Time
	FakeSchedule []*time.Time

	Parent     ComplicatedPerson
	FakeParent *ComplicatedPerson

	Children     []ComplicatedPerson
	FakeChildren []*ComplicatedPerson
}

func NewComplicatedMan() ComplicatedMan {
	numSche := rand.Intn(10)
	numFSche := rand.Intn(10)
	numC := rand.Intn(10)
	numFC := rand.Intn(10)

	sche := make([]time.Time, 0, numSche)
	fSche := make([]*time.Time, 0, numFSche)
	c := make([]ComplicatedPerson, 0, numC)
	fc := make([]*ComplicatedPerson, 0, numFC)

	for i := 0; i < numSche; i++ {
		sche = append(sche, randTime())
	}

	for i := 0; i < numFSche; i++ {
		fSche = append(fSche, randTimePtr())
	}

	for i := 0; i < numC; i++ {
		c = append(c, NewComplicatedPerson())
	}

	for i := 0; i < numFC; i++ {
		fc = append(fc, NewComplicatedPersonPtr())
	}

	return ComplicatedMan{
		Name:         "name:" + randStr(),
		FakeName:     randStrPtr(),
		Age:          randInt(),
		FakeAge:      randIntPtr(),
		Birth:        randTime(),
		FakeBirth:    randTimePtr(),
		Schedule:     sche,
		FakeSchedule: fSche,
		Parent:       NewComplicatedPerson(),
		FakeParent:   NewComplicatedPersonPtr(),
		Children:     c,
		FakeChildren: fc,
	}
}

type ComplicatedPerson struct {
	Name     string
	FakeName *string

	Age     int
	FakeAge *int

	Birth     time.Time
	FakeBirth *time.Time

	Schedule     []time.Time
	FakeSchedule []*time.Time

	Parent     Person
	FakeParent *Person

	Children     []Person
	FakeChildren []*Person
}

func NewComplicatedPerson() ComplicatedPerson {
	numSche := rand.Intn(10)
	numFSche := rand.Intn(10)
	numC := rand.Intn(10)
	numFC := rand.Intn(10)

	sche := make([]time.Time, 0, numSche)
	fSche := make([]*time.Time, 0, numFSche)
	c := make([]Person, 0, numC)
	fc := make([]*Person, 0, numFC)

	for i := 0; i < numSche; i++ {
		sche = append(sche, randTime())
	}

	for i := 0; i < numFSche; i++ {
		fSche = append(fSche, randTimePtr())
	}

	for i := 0; i < numC; i++ {
		c = append(c, NewPerson())
	}

	for i := 0; i < numFC; i++ {
		fc = append(fc, NewPersonPtr())
	}

	return ComplicatedPerson{
		Name:         randStr(),
		FakeName:     randStrPtr(),
		Age:          randInt(),
		FakeAge:      randIntPtr(),
		Birth:        randTime(),
		FakeBirth:    randTimePtr(),
		Schedule:     sche,
		FakeSchedule: fSche,
		Parent:       NewPerson(),
		FakeParent:   NewPersonPtr(),
		Children:     c,
		FakeChildren: fc,
	}
}

func NewComplicatedPersonPtr() *ComplicatedPerson {
	p := NewComplicatedPerson()
	return &p
}

type Person struct {
	Name     string
	FakeName *string

	Age     int
	FakeAge *int

	Birth     time.Time
	FakeBirth *time.Time

	Schedule     []time.Time
	FakeSchedule []*time.Time
}

func NewPerson() Person {
	numSche := rand.Intn(10)
	numFSche := rand.Intn(10)

	sche := make([]time.Time, 0, numSche)
	fSche := make([]*time.Time, 0, numFSche)

	for i := 0; i < numSche; i++ {
		sche = append(sche, randTime())
	}

	for i := 0; i < numFSche; i++ {
		fSche = append(fSche, randTimePtr())
	}

	return Person{
		Name:         randStr(),
		FakeName:     randStrPtr(),
		Age:          randInt(),
		FakeAge:      randIntPtr(),
		Birth:        randTime(),
		FakeBirth:    randTimePtr(),
		Schedule:     sche,
		FakeSchedule: fSche,
	}
}

func NewPersonPtr() *Person {
	p := NewPerson()
	return &p
}

func randStr() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func randStrPtr() *string {
	s := randStr()
	return &s
}

func randInt() int {
	return rand.Int()
}

func randIntPtr() *int {
	i := randInt()
	return &i
}

func randTime() time.Time {
	t := time.Date(rand.Intn(1000)+2000, time.Month(rand.Intn(12)+1), rand.Intn(30)+1, rand.Intn(24), rand.Intn(60), rand.Intn(60), rand.Int(), time.UTC)
	return t
}

func randTimePtr() *time.Time {
	t := randTime()
	return &t
}
