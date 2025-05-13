// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ackone/cluster"
	membershipattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ackone/membershipattachment"
	addresspool "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/customline"
	domain "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/gtminstance"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/record"
	domaincdn "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domainconfig"
	fctrigger "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/fctrigger"
	alarmcontactgroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cloudmonitorservice/alarmcontactgroup"
	autoscalingconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/autoscalingconfig"
	edgekubernetes "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/edgekubernetes"
	kubernetes "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/kubernetes"
	kubernetesaddon "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/kubernetesaddon"
	kubernetesnodepool "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/kubernetesnodepool"
	kubernetespermissions "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/kubernetespermissions"
	managedkubernetes "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/managedkubernetes"
	serverlesskubernetes "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cs/serverlesskubernetes"
	activation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/activation"
	autoprovisioninggroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autoprovisioninggroup"
	autosnapshotpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autosnapshotpolicy"
	autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autosnapshotpolicyattachment"
	capacityreservation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/capacityreservation"
	command "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/command"
	dedicatedhost "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/dedicatedhost"
	dedicatedhostcluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/dedicatedhostcluster"
	deploymentset "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/deploymentset"
	disk "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/disk"
	diskattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/diskattachment"
	elasticityassurance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/elasticityassurance"
	hpccluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/hpccluster"
	image "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/image"
	imagecomponent "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagecomponent"
	imagecopy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagecopy"
	imageexport "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imageexport"
	imageimport "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imageimport"
	imagepipeline "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagepipeline"
	imagepipelineexecution "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagepipelineexecution"
	imagesharepermission "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagesharepermission"
	instanceecs "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/instance"
	instanceset "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/instanceset"
	invocation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/invocation"
	keypair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/keypair"
	keypairattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/keypairattachment"
	launchtemplate "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterfaceattachment"
	networkinterfacepermission "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterfacepermission"
	prefixlist "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/prefixlist"
	reservedinstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/reservedinstance"
	securitygroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/securitygrouprule"
	sessionmanagerstatus "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/sessionmanagerstatus"
	snapshot "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/snapshot"
	snapshotgroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/snapshotgroup"
	storagecapacityunit "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/storagecapacityunit"
	alias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/alias"
	instancekms "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/instance"
	key "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/key"
	secret "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/secret"
	endpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpoint"
	endpointacl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpointacl"
	queue "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/queue"
	subscription "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/subscription"
	topic "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/topic"
	accesspoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/accountpublicaccessblock"
	bucket "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucket"
	bucketaccessmonitor "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketaccessmonitor"
	bucketacl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketacl"
	bucketcname "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcname"
	bucketcnametoken "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcnametoken"
	bucketcors "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcors"
	bucketdataredundancytransition "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketdataredundancytransition"
	buckethttpsconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/buckethttpsconfig"
	bucketlogging "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketlogging"
	bucketmetaquery "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketmetaquery"
	bucketobject "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketpublicaccessblock"
	bucketreferer "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketreferer"
	bucketreplication "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketreplication"
	bucketrequestpayment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketrequestpayment"
	bucketserversideencryption "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketserversideencryption"
	bucketstyle "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketstyle"
	buckettransferacceleration "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/buckettransferacceleration"
	bucketuserdefinedlogfields "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketuserdefinedlogfields"
	bucketversioning "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketversioning"
	bucketwebsite "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketwebsite"
	bucketworm "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketworm"
	account "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/account"
	accountprivilege "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/backuppolicy"
	clusterpolardb "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/clusterendpoint"
	database "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/database"
	endpointpolardb "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpointaddress"
	globaldatabasenetwork "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/globaldatabasenetwork"
	parametergroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/parametergroup"
	primaryendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/primaryendpoint"
	providerconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/providerconfig"
	accesskey "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accesskey"
	accountalias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accountpasswordpolicy"
	group "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/group"
	groupmembership "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/grouppolicyattachment"
	loginprofile "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/loginprofile"
	passwordpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/passwordpolicy"
	policy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/policy"
	role "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/samlprovider"
	securitypreference "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/securitypreference"
	user "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/user"
	usergroupattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/usergroupattachment"
	userpolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/userpolicyattachment"
	accounttair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/account"
	auditlogconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/auditlogconfig"
	connection "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/connection"
	instancetair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/instance"
	tairinstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/tairinstance"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		membershipattachment.Setup,
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instance.Setup,
		monitorconfig.Setup,
		record.Setup,
		domaincdn.Setup,
		domainconfig.Setup,
		fctrigger.Setup,
		alarmcontactgroup.Setup,
		autoscalingconfig.Setup,
		edgekubernetes.Setup,
		kubernetes.Setup,
		kubernetesaddon.Setup,
		kubernetesnodepool.Setup,
		kubernetespermissions.Setup,
		managedkubernetes.Setup,
		serverlesskubernetes.Setup,
		activation.Setup,
		autoprovisioninggroup.Setup,
		autosnapshotpolicy.Setup,
		autosnapshotpolicyattachment.Setup,
		capacityreservation.Setup,
		command.Setup,
		dedicatedhost.Setup,
		dedicatedhostcluster.Setup,
		deploymentset.Setup,
		disk.Setup,
		diskattachment.Setup,
		elasticityassurance.Setup,
		hpccluster.Setup,
		image.Setup,
		imagecomponent.Setup,
		imagecopy.Setup,
		imageexport.Setup,
		imageimport.Setup,
		imagepipeline.Setup,
		imagepipelineexecution.Setup,
		imagesharepermission.Setup,
		instanceecs.Setup,
		instanceset.Setup,
		invocation.Setup,
		keypair.Setup,
		keypairattachment.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacepermission.Setup,
		prefixlist.Setup,
		reservedinstance.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		sessionmanagerstatus.Setup,
		snapshot.Setup,
		snapshotgroup.Setup,
		storagecapacityunit.Setup,
		alias.Setup,
		instancekms.Setup,
		key.Setup,
		secret.Setup,
		endpoint.Setup,
		endpointacl.Setup,
		queue.Setup,
		subscription.Setup,
		topic.Setup,
		accesspoint.Setup,
		accountpublicaccessblock.Setup,
		bucket.Setup,
		bucketaccessmonitor.Setup,
		bucketacl.Setup,
		bucketcname.Setup,
		bucketcnametoken.Setup,
		bucketcors.Setup,
		bucketdataredundancytransition.Setup,
		buckethttpsconfig.Setup,
		bucketlogging.Setup,
		bucketmetaquery.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreferer.Setup,
		bucketreplication.Setup,
		bucketrequestpayment.Setup,
		bucketserversideencryption.Setup,
		bucketstyle.Setup,
		buckettransferacceleration.Setup,
		bucketuserdefinedlogfields.Setup,
		bucketversioning.Setup,
		bucketwebsite.Setup,
		bucketworm.Setup,
		account.Setup,
		accountprivilege.Setup,
		backuppolicy.Setup,
		clusterpolardb.Setup,
		clusterendpoint.Setup,
		database.Setup,
		endpointpolardb.Setup,
		endpointaddress.Setup,
		globaldatabasenetwork.Setup,
		parametergroup.Setup,
		primaryendpoint.Setup,
		providerconfig.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		loginprofile.Setup,
		passwordpolicy.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		securitypreference.Setup,
		user.Setup,
		usergroupattachment.Setup,
		userpolicyattachment.Setup,
		accounttair.Setup,
		auditlogconfig.Setup,
		connection.Setup,
		instancetair.Setup,
		tairinstance.Setup,
		vpc.Setup,
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
