// package handlers

// import (
// 	"context"
// 	"grpcapi/internals/models"
// 	"grpcapi/internals/repositories/mongodb"
// 	pb "grpcapi/proto/gen"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	for _, teacher := range req.GetTeachers() {
// 		if teacher.Id != "" {
// 			return nil, status.Error(codes.InvalidArgument, "request is in incorrect format: non-empty ID fields are not allowed.")
// 		}
// 	}

// 	addedTeachers, err := mongodb.AddTeachersToDb(ctx, req.GetTeachers())
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}
// 	return &pb.Teachers{Teachers: addedTeachers}, nil
// }

// func (s *Server) GetTeachers(ctx context.Context, req *pb.GetTeachersRequest) (*pb.Teachers, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	// Filtering, getting the filters from the request, another function
// 	filter, err := buildFilter(req.Teacher, &models.Teacher{})
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	// Sorting, getting the sort options from the request, another function
// 	sortOptions := buildSortOptions(req.GetSortBy())
// 	// Access the database to fetch data, another function

// 	teachers, err := mongodb.GetTeachersFromDb(ctx, sortOptions, filter)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.Teachers{Teachers: teachers}, nil
// }

// func (s *Server) UpdateTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	updatedTeachers, err := mongodb.ModifyTeachersInDb(ctx, req.Teachers)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.Teachers{Teachers: updatedTeachers}, nil
// }

// func (s *Server) DeleteTeachers(ctx context.Context, req *pb.TeacherIds) (*pb.DeleteTeachersConfirmation, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	ids := req.GetIds()
// 	var teacherIdsToDelete []string
// 	for _, teacher := range ids {
// 		teacherIdsToDelete = append(teacherIdsToDelete, teacher.Id)
// 	}

// 	deletedIds, err := mongodb.DeleteTeachersFromDb(ctx, teacherIdsToDelete)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.DeleteTeachersConfirmation{
// 		Status:     "Teachers successfully deleted",
// 		DeletedIds: deletedIds,
// 	}, nil
// }

// func (s *Server) GetStudentsByClassTeacher(ctx context.Context, req *pb.TeacherId) (*pb.Students, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	teacherId := req.GetId()

// 	students, err := mongodb.GetStudentsByTeacherIdFromDb(ctx, teacherId)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.Students{Students: students}, nil
// }

// func (s *Server) GetStudentCountByClassTeacher(ctx context.Context, req *pb.TeacherId) (*pb.StudentCount, error) {
// 	err := req.Validate()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	teacherId := req.GetId()

// 	count, err := mongodb.GetStudentCountByTeacherIdFromDb(ctx, teacherId)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &pb.StudentCount{Status: true, StudentCount: int32(count)}, nil
// }
package handlers

import (
	"context"
	"grpcapi/internals/repositories/mongodb"
	"grpcapi/internals/models" 
	"grpcapi/pkg/utils"
	pb "grpcapi/proto/gen"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddTeachers adds multiple teachers to MongoDB
func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
	// Create MongoDB client
	client, err := mongodb.CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "internal error")
	}
	defer client.Disconnect(ctx)

	// Slice to hold model teachers
	newTeachers := make([]*models.Teacher, len(req.GetTeachers()))

	// Convert pb.Teacher → models.Teacher using reflection
	for i, pbTeacher := range req.GetTeachers() {
		modelTeacher := models.Teacher{}

		pbVal := reflect.ValueOf(pbTeacher).Elem()
		modelVal := reflect.ValueOf(&modelTeacher).Elem()

		for j := 0; j < pbVal.NumField(); j++ {
			pbField := pbVal.Field(j)
			fieldName := pbVal.Type().Field(j).Name
			modelField := modelVal.FieldByName(fieldName)

			if modelField.IsValid() && modelField.CanSet() {
				modelField.Set(pbField)
			}
		}

		newTeachers[i] = &modelTeacher
	}

	// Insert teachers into MongoDB
	addedTeachers := []*pb.Teacher{}
	for _, teacher := range newTeachers {
		result, err := client.Database("school").Collection("teachers").InsertOne(ctx, teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error adding teacher to database")
		}

		if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
			teacher.Id = objectId.Hex()
		}

		// Convert models.Teacher → pb.Teacher
		pbTeacher := &pb.Teacher{}
		modelVal := reflect.ValueOf(*teacher)
		pbVal := reflect.ValueOf(pbTeacher).Elem()

		for j := 0; j < modelVal.NumField(); j++ {
			modelField := modelVal.Field(j)
			modelFieldType := modelVal.Type().Field(j)
			pbField := pbVal.FieldByName(modelFieldType.Name)

			if pbField.IsValid() && pbField.CanSet() {
				pbField.Set(modelField)
			}
		}

		addedTeachers = append(addedTeachers, pbTeacher)
	}

	// Return final response
	return &pb.Teachers{Teachers: addedTeachers}, nil
}
