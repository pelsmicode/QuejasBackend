UPDATE companies
	SET name=$1,
	township_id=$2,
	nit=$3,
	address=$4,
	phone=$5,
	email=$6
	WHERE id=$7
