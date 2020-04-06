package middle

import "testing"

func TestAbc(t *testing.T) {
	c := Constructor()
	c.AddNum(1)
	c.AddNum(2)
	c.AddNum(6)
	c.AddNum(4)
	c.FindMedian()
}
