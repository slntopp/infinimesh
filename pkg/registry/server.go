//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package registry

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infinimesh/infinimesh/pkg/node"
	"github.com/infinimesh/infinimesh/pkg/node/dgraph"
	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
	"github.com/infinimesh/infinimesh/pkg/registry/registrypb"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dgraph-io/dgo"
)

//Server is a Data type for Device Controller file
type Server struct {
	dgo *dgo.Dgraph

	repo node.Repo
	Log  *zap.Logger
}

//NewServer is a method to create the Dgraph Server for Device registry
func NewServer(dg *dgo.Dgraph) *Server {
	return &Server{
		dgo: dg,
		repo: &dgraph.DGraphRepo{
			Dg: dg,
		},
	}
}

var a node.AccountController

func (s *Server) getFingerprint(pemCert []byte, certType string) (fingerprint []byte, err error) {
	pemBlock, _ := pem.Decode(pemCert)
	if pemBlock == nil {
		return nil, errors.New("Could not decode PEM")
	}
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return sha256Sum(cert.Raw), nil
}

func sha256Sum(c []byte) []byte {
	s := sha256.New()
	_, err := s.Write(c)
	if err != nil {
		panic(err)
	}
	return s.Sum(nil)
}

//Create is a method for creating Devices
func (s *Server) Create(ctx context.Context, request *registrypb.CreateRequest) (*registrypb.CreateResponse, error) {

	log := s.Log.Named("Create Device Controller")

	//Added logging
	log.Info("Function Invoked",
		zap.String("Device Name", request.Device.Name),
		zap.String("Namespace", request.Device.Namespace),
		zap.Bool("Enabled", request.Device.Enabled.Value),
	)

	_, err := s.repo.GetNamespaceID(ctx, request.Device.Namespace)
	if err != nil {
		log.Error("Data Validation for Device Creation", zap.String("Error", "The Namespace cannot not be empty"))
		return nil, status.Error(codes.FailedPrecondition, "The Namespace cannot not be empty")
	}

	if request.Device.Certificate == nil {
		log.Error("Data Validation for Device Creation", zap.String("Error", "No certificate provided"))
		return nil, status.Error(codes.FailedPrecondition, "No certificate provided")
	}

	//Initialize the Account Controller with Device controller data
	a.Repo = s.repo
	a.Log = s.Log

	//Get metadata from context and perform validation
	_, requestorID, err := node.Validation(ctx, log)
	if err != nil {
		return nil, err
	}

	//Check if the user has access to create the device for the namespace
	authresp, err := a.IsAuthorizedNamespace(ctx, &nodepb.IsAuthorizedNamespaceRequest{
		Namespaceid: request.Device.Namespace,
		Account:     requestorID,
		Action:      nodepb.Action_WRITE,
	})
	if err != nil {
		log.Error("Unable to get permissions for the account", zap.Error(err))
		return nil, status.Error(codes.PermissionDenied, "Unable to get permissions for the account")
	}
	if !authresp.GetDecision().GetValue() {
		log.Error("The Account does not have permission to create Device")
		return nil, status.Error(codes.PermissionDenied, "The account does not have permission to create device")
	}

	resp, err := s.CreateQ(ctx, request)
	if err != nil {
		//Added logging
		log.Error("Failed to create Device", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	//Added logging
	log.Info("Device Created", zap.String("Device ID", resp.Device.Id), zap.String("Device Name", resp.Device.Name))
	return resp, nil
}

//Update is a method for updating Devices details
func (s *Server) Update(ctx context.Context, request *registrypb.UpdateRequest) (response *registrypb.UpdateResponse, err error) {

	log := s.Log.Named("Update Device Controller")

	//Added logging
	log.Info("Function Invoked",
		zap.String("Device", request.Device.Id),
		zap.Any("FieldMask", request.GetFieldMask()),
	)

	//Initialize the Account Controller with Device controller data
	a.Repo = s.repo
	a.Log = s.Log

	//Get metadata from context and perform validation
	_, requestorID, err := node.Validation(ctx, log)
	if err != nil {
		return nil, err
	}

	//Check if the user has access to update the device
	authresp, err := a.IsAuthorized(ctx, &nodepb.IsAuthorizedRequest{
		Node:    request.Device.Id,
		Account: requestorID,
		Action:  nodepb.Action_WRITE,
	})
	if err != nil {
		log.Error("Unable to get permissions for the account", zap.Error(err))
		return nil, status.Error(codes.PermissionDenied, "Unable to get permissions for the account")
	}
	if !authresp.GetDecision().GetValue() {
		log.Error("The Account does not have permission to create Device")
		return nil, status.Error(codes.PermissionDenied, "The account does not have permission to create device")
	}

	resp, err := s.UpdateQ(ctx, request)
	if err != nil {
		//Added logging
		log.Error("Failed to update Device", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	//Added logging
	log.Info("Device successfully updated")
	return resp, nil
}

//GetByFingerprint is a method for get FringerPrint for a Device
func (s *Server) GetByFingerprint(ctx context.Context, request *registrypb.GetByFingerprintRequest) (*registrypb.GetByFingerprintResponse, error) {

	resp, err := s.GetByFingerprintQ(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return resp, nil
}

//Get is a method for get details for a Device
func (s *Server) Get(ctx context.Context, request *registrypb.GetRequest) (response *registrypb.GetResponse, err error) {

	log := s.Log.Named("Get Device Controller")

	//Added logging
	log.Info("Function Invoked", zap.String("Device", request.Id))

	//Initialize the Account Controller with Device controller data
	a.Repo = s.repo
	a.Log = s.Log

	//Get metadata from context and perform validation
	_, requestorID, err := node.Validation(ctx, log)
	if err != nil {
		return nil, err
	}

	//Check if the user has access to get the device details
	authresp, err := a.IsAuthorized(ctx, &nodepb.IsAuthorizedRequest{
		Node:    request.Id,
		Account: requestorID,
		Action:  nodepb.Action_READ,
	})
	if err != nil {
		log.Error("Unable to get permissions for the account", zap.Error(err))
		return nil, status.Error(codes.PermissionDenied, "Unable to get permissions for the account")
	}

	if !authresp.GetDecision().GetValue() {
		log.Error("The Account does not have permission to get Device list")
		return nil, status.Error(codes.PermissionDenied, "The Account does not have permission to get Device list")
	}

	resp, err := s.GetQ(ctx, request)
	if err != nil {
		log.Error("Failed to Get Device", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info("Get Device details successsful")
	return resp, nil
}

//List is a method that list all Devices for a specific Namespace
func (s *Server) List(ctx context.Context, request *registrypb.ListDevicesRequest) (response *registrypb.ListResponse, err error) {

	log := s.Log.Named("List Device Controller")

	//Added logging
	log.Info("Function Invoked", zap.String("Account", request.Account), zap.String("Namespace", request.Namespace))

	//Initialize the Account Controller with Device controller data
	a.Repo = s.repo
	a.Log = s.Log

	//Get metadata from context and perform validation
	_, requestorID, err := node.Validation(ctx, log)
	if err != nil {
		return nil, err
	}

	isRoot, err := a.IsRoot(ctx, &nodepb.IsRootRequest{
		Account: requestorID,
	})
	if err != nil {
		return nil, err
	}

	//If Account is root provide all access
	if isRoot.IsRoot {
		return s.List(ctx, &registrypb.ListDevicesRequest{Namespace: request.Namespace})
	}

	//Check if the non root user has access to the namespace
	authresp, err := a.IsAuthorizedNamespace(ctx, &nodepb.IsAuthorizedNamespaceRequest{
		Namespaceid: request.Namespace,
		Account:     requestorID,
		Action:      nodepb.Action_READ,
	})
	if err != nil {
		log.Error("Unable to get permissions for the account", zap.Error(err))
		return nil, status.Error(codes.PermissionDenied, "Unable to get permissions for the account")
	}
	if !authresp.GetDecision().GetValue() {
		log.Error("The Account does not have permission to list Devices")
		return nil, status.Error(codes.PermissionDenied, "The account does not have permission to list Devices")
	}

	resp, err := s.ListForAccountQ(ctx, request)
	if err != nil {
		log.Error("Failed to list Devices", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info("List Devices successsful")
	return resp, nil
}

//Delete is a method that deletes a Device
func (s *Server) Delete(ctx context.Context, request *registrypb.DeleteRequest) (response *registrypb.DeleteResponse, err error) {

	log := s.Log.Named("Delete Device Controller")

	//Added logging
	log.Info("Function Invoked", zap.String("Device", request.Id))

	//Initialize the Account Controller with Device controller data
	a.Repo = s.repo
	a.Log = s.Log

	//Get metadata from context and perform validation
	_, requestorID, err := node.Validation(ctx, log)
	if err != nil {
		return nil, err
	}

	//Check if the user has access to delete the device
	authresp, err := a.IsAuthorized(ctx, &nodepb.IsAuthorizedRequest{
		Node:    request.Id,
		Account: requestorID,
		Action:  nodepb.Action_WRITE,
	})
	if err != nil {
		log.Error("Unable to get permissions for the account", zap.Error(err))
		return nil, status.Error(codes.PermissionDenied, "Unable to get permissions for the account")
	}

	if !authresp.GetDecision().GetValue() {
		log.Error("The Account does not have permission to delete the device")
		return nil, status.Error(codes.PermissionDenied, "The Account does not have permission to delete the device")
	}

	resp, err := s.DeleteQ(ctx, request)
	if err != nil {
		log.Error("Failed to delete Device", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info("Devices deleted successsful")
	return resp, nil
}
