package job

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateJob(ctx context.Context, job *Job) (uuid.UUID, error)
	GetJob(ctx context.Context, id uuid.UUID) (*Job, error)
	UpdateStatus(context.Context, uuid.UUID, string) error
}

type JobStore struct {
	db *pgxpool.Pool
}

func NewJob(db *pgxpool.Pool) *JobStore {
	return &JobStore{
		db: db,
	}
}

func (j *JobStore) CreateJob(ctx context.Context, jj *Job) (uuid.UUID, error) {
	const query = `
	insert into jobs (
	id,
	payload,
	) values(
	 $1,
	 $2,
	 $3
	 )
	 returning id;
	`
	var id uuid.UUID
	row := j.db.QueryRow(ctx, query, id, jj.Payload, 1)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (j *JobStore) GetJob(ctx context.Context, jobID uuid.UUID) (*Job, error) {
	const query = `
		select
			id,
			payload,
			status
		from jobs
		where id = $1
		`
	job := new(Job)
	if err := j.db.QueryRow(ctx, query, jobID).Scan(&job.ID, &job.Payload, &job.Status); err != nil {
		return nil, err
	}
	return job, nil
}

func (j *JobStore) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	const query = `
	UPDATE JOBS SET STATUS = $1 where id = $2
	`
	_, err := j.db.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}
	return nil

}
