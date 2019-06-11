package demo1

import "fmt"

type Job interface {
	Run()
}

type JobOne struct {

}

func (job *JobOne)Run()   {
	fmt.Println("hello Job1")
}

type JobTwo struct {

}

func (job *JobTwo)Run()  {
	fmt.Println("hello Job2")
}

var(
	 sch map[string]Job
)

func main()  {
	sch = make(map[string]Job)

	sch["jb1"] = &JobOne{}
	sch["jb2"] = &JobTwo{}

	for v,k := range sch {
		k.Run()
		fmt.Println(v)
	}
}