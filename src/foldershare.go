package foldershare

type FileInfo struct {
	name  string
	cdate FileTime
	mdate FileTime
}

//Equal
func Equal(origin, my FileInfo) bool {
	return origin.name == my.name && origin.cdate == my.cdate && origin.mdate == my.mdate
}

func (self *FileInfo) IsNewer(other FileInfo) bool {
	if self.mdate.IsNewer(other.mdate) {
		return true
	}
	if self.cdate.IsNewer(other.cdate) {
		return true
	}

	return false
}

