package service

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"sync"
)

func (s *Service) GetJobPostingByIDService(ctx context.Context, jid uint64) (models.Jobs, error) {
	if jid < uint64(0) {
		return models.Jobs{}, errors.New("id cannot be 0")
	}
	jobData, err := s.UserRepo.FetchJobPostingByID(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) GetAllJobPostingsService(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FetchAllJobPostings(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil
}

func (s *Service) CreateJobPostingService(ctx context.Context, jobData models.NewJobRequest, cid uint64) (models.NewJobResponse, error) {
	jobData.Cid = uint(cid)
	jobDatas, err := s.UserRepo.InsertJobPosting(ctx, jobData)
	if err != nil {
		return models.NewJobResponse{}, err
	}
	return jobDatas, nil
}

func (s *Service) ListJobsForCompanyService(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FetchJobsForCompany(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (s *Service) ApplicationProcessor(ctx context.Context, jobApplications []models.RequestJob) ([]models.RequestJob, error) {
	ch := make(chan models.RequestJob)
	var wg sync.WaitGroup

	for _, application := range jobApplications {
		wg.Add(1)

		application := application
		go func(app models.RequestJob) {
			defer wg.Done()
			job, err := s.UserRepo.FetchJobPostingByID(ctx, application.Jid)
			if err != nil {
				return
			}
			if validateJobApplication(app, job) {
				ch <- app
			}
		}(application)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := []models.RequestJob{}
	for app := range ch {
		result = append(result, app)
	}

	return result, nil
}

func validateJobApplication(application models.RequestJob, job models.Jobs) bool {

	if application.MinNP < job.MinNP || application.MaxNP > job.MaxNP {
		return false
	}

	if application.Budget < job.Budget {
		return false
	}

	/*if !containsAll(job.Locations, application.LocationsIDs) {
		return false
	}

	if !containsAll(job.TechnologyStack, application.TechnologyStackIDs) {
		return false
	}

	if !containsAll(job.WorkMode, application.WorkModeIDs) {
		return false
	}*/

	if application.Description != job.Description {
		return false
	}

	if application.MinExp < job.MinExp || application.MaxExp > job.MaxExp {
		return false
	}

	//if !containsAll(job.Qualification, application.Qualification) {
	//	return false
	//}
	//
	//if !containsAll(job.Shift, application.Shift) {
	//	return false
	//}
	//
	//if !containsAll(job.JobType, application.JobType) {
	//	return false
	//}

	return true
}
