# Start

$ docker run -d \
	--name mypaas \
	-e POSTGRES_USER=mypaas \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v $(pwd)/db:/var/lib/postgresql/data \
    -p 5432:5432 \
	postgres