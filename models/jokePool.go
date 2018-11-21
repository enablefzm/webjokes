package models

var OBJokePool *JokePool = NewJokePool()

type JokePool struct {
	arrJoke []*JokeSource
}

func NewJokePool() *JokePool {
	p := &JokePool{
		arrJoke: make([]*JokeSource, 0, 30),
	}
	return p
}

func (this *JokePool) loadJokes() int {
	rss, err := DBSave.QuerysLimit("*", "joke_text", "is_check=0", 1, 30, "id DESC")
	if err != nil {
		return 0
	}
	il := len(rss)
	if il < 1 {
		return il
	}
	this.arrJoke = make([]*JokeSource, 0, 30)
	for _, rs := range rss {
		this.arrJoke = append(this.arrJoke, NewJokeSourceOnRs(rs))
	}
	return il
}

func (this *JokePool) Get() (*JokeSource, bool) {
	ptJoke, ok := this.getJoke()
	if !ok {
		// 重新装载对象
		resValue := this.loadJokes()
		if resValue < 1 {
			return nil, false
		}
		ptJoke, ok = this.getJoke()
	}
	return ptJoke, ok
}

func (this *JokePool) getJoke() (*JokeSource, bool) {
	var k int
	var ptJoke *JokeSource
	for k, ptJoke = range this.arrJoke {
		if ptJoke != nil {
			this.arrJoke[k] = nil
			return ptJoke, true
		}
	}
	return nil, false
}

func (this *JokePool) Count() int {
	i := 0
	for _, p := range this.arrJoke {
		if p != nil {
			i++
		}
	}
	return i
}

//	func init() {
//		OBJokePool = NewJokePool()
//	}
