package domain


//type Alias1 = string
type list struct {
	Program_name int
	OrderByDate_Desc int
	OrderByDate_Asc int
	OrderByLocation int
}

var Enum = &list{
	OrderByDate_Desc: 1,
	OrderByDate_Asc:  2,
	OrderByLocation:  3,
	Program_name: 4,
}


