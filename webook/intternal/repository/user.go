package repository

type UserRepository struct {
}

func (r *UserRepository) FindById(int64) {
	//先从cache中找
	//再从dao中找
	//找到了回写cache
}
