// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./accountProject";
export * from "./accountSshKey";
export * from "./appleSiliconServer";
export * from "./baremetalServer";
export * from "./cockpit";
export * from "./cockpitGrafanaUser";
export * from "./cockpitToken";
export * from "./container";
export * from "./containerCron";
export * from "./containerDomain";
export * from "./containerNamespace";
export * from "./containerToken";
export * from "./domainRecord";
export * from "./domainZone";
export * from "./flexibleIp";
export * from "./function";
export * from "./functionCron";
export * from "./functionDomain";
export * from "./functionNamespace";
export * from "./functionToken";
export * from "./getAccountProject";
export * from "./getAccountSshKey";
export * from "./getBaremetalOffer";
export * from "./getBaremetalOption";
export * from "./getBaremetalOs";
export * from "./getBaremetalServer";
export * from "./getCockpit";
export * from "./getContainer";
export * from "./getContainerNamespace";
export * from "./getDomainRecord";
export * from "./getDomainZone";
export * from "./getFlexibleIp";
export * from "./getFunction";
export * from "./getFunctionNamespace";
export * from "./getIamApplication";
export * from "./getIamGroup";
export * from "./getIamSshKey";
export * from "./getIamUser";
export * from "./getInstanceImage";
export * from "./getInstanceIp";
export * from "./getInstancePrivateNic";
export * from "./getInstanceSecurityGroup";
export * from "./getInstanceServer";
export * from "./getInstanceServers";
export * from "./getInstanceSnapshot";
export * from "./getInstanceVolume";
export * from "./getIotDevice";
export * from "./getIotHub";
export * from "./getK8sCluster";
export * from "./getK8sPool";
export * from "./getK8sVersion";
export * from "./getLb";
export * from "./getLbAcls";
export * from "./getLbBackend";
export * from "./getLbBackends";
export * from "./getLbCertificate";
export * from "./getLbFrontend";
export * from "./getLbFrontends";
export * from "./getLbIp";
export * from "./getLbIps";
export * from "./getLbRoute";
export * from "./getLbRoutes";
export * from "./getLbs";
export * from "./getMarketplaceImage";
export * from "./getObjectBucket";
export * from "./getRdbAcl";
export * from "./getRdbDatabase";
export * from "./getRdbDatabaseBackup";
export * from "./getRdbInstance";
export * from "./getRdbPrivilege";
export * from "./getRedisCluster";
export * from "./getRegistryImage";
export * from "./getRegistryNamespace";
export * from "./getSecret";
export * from "./getSecretVersion";
export * from "./getTemDomain";
export * from "./getVpcGatewayNetwork";
export * from "./getVpcPrivateNetwork";
export * from "./getVpcPublicGateway";
export * from "./getVpcPublicGatewayDhcp";
export * from "./getVpcPublicGatewayDhcpReservation";
export * from "./getVpcPublicGatewayIp";
export * from "./getVpcPublicGatewayPatRule";
export * from "./getWebhostingOffer";
export * from "./iamApiKey";
export * from "./iamApplication";
export * from "./iamGroup";
export * from "./iamPolicy";
export * from "./iamSshKey";
export * from "./instanceImage";
export * from "./instanceIp";
export * from "./instanceIpReverseDns";
export * from "./instancePlacementGroup";
export * from "./instancePrivateNic";
export * from "./instanceSecurityGroup";
export * from "./instanceSecurityGroupRules";
export * from "./instanceServer";
export * from "./instanceSnapshot";
export * from "./instanceUserData";
export * from "./instanceVolume";
export * from "./iotDevice";
export * from "./iotHub";
export * from "./iotNetwork";
export * from "./iotRoute";
export * from "./k8sCluster";
export * from "./k8sPool";
export * from "./lb";
export * from "./lbBackend";
export * from "./lbCertificate";
export * from "./lbFrontend";
export * from "./lbIp";
export * from "./lbRoute";
export * from "./mnqCredential";
export * from "./mnqNamespace";
export * from "./objectBucket";
export * from "./objectBucketAcl";
export * from "./objectBucketLockConfiguration";
export * from "./objectBucketPolicy";
export * from "./objectBucketWebsiteConfiguration";
export * from "./objectItem";
export * from "./provider";
export * from "./rdbAcl";
export * from "./rdbDatabase";
export * from "./rdbDatabaseBackup";
export * from "./rdbInstance";
export * from "./rdbPrivilege";
export * from "./rdbReadReplica";
export * from "./rdbUser";
export * from "./redisCluster";
export * from "./registryNamespace";
export * from "./secret";
export * from "./secretVersion";
export * from "./temDomain";
export * from "./vpcGatewayNetwork";
export * from "./vpcPrivateNetwork";
export * from "./vpcPublicGateway";
export * from "./vpcPublicGatewayDhcp";
export * from "./vpcPublicGatewayDhcpReservation";
export * from "./vpcPublicGatewayIp";
export * from "./vpcPublicGatewayIpReverseDns";
export * from "./vpcPublicGatewayPatRule";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { AccountProject } from "./accountProject";
import { AccountSshKey } from "./accountSshKey";
import { AppleSiliconServer } from "./appleSiliconServer";
import { BaremetalServer } from "./baremetalServer";
import { Cockpit } from "./cockpit";
import { CockpitGrafanaUser } from "./cockpitGrafanaUser";
import { CockpitToken } from "./cockpitToken";
import { Container } from "./container";
import { ContainerCron } from "./containerCron";
import { ContainerDomain } from "./containerDomain";
import { ContainerNamespace } from "./containerNamespace";
import { ContainerToken } from "./containerToken";
import { DomainRecord } from "./domainRecord";
import { DomainZone } from "./domainZone";
import { FlexibleIp } from "./flexibleIp";
import { Function } from "./function";
import { FunctionCron } from "./functionCron";
import { FunctionDomain } from "./functionDomain";
import { FunctionNamespace } from "./functionNamespace";
import { FunctionToken } from "./functionToken";
import { IamApiKey } from "./iamApiKey";
import { IamApplication } from "./iamApplication";
import { IamGroup } from "./iamGroup";
import { IamPolicy } from "./iamPolicy";
import { IamSshKey } from "./iamSshKey";
import { InstanceImage } from "./instanceImage";
import { InstanceIp } from "./instanceIp";
import { InstanceIpReverseDns } from "./instanceIpReverseDns";
import { InstancePlacementGroup } from "./instancePlacementGroup";
import { InstancePrivateNic } from "./instancePrivateNic";
import { InstanceSecurityGroup } from "./instanceSecurityGroup";
import { InstanceSecurityGroupRules } from "./instanceSecurityGroupRules";
import { InstanceServer } from "./instanceServer";
import { InstanceSnapshot } from "./instanceSnapshot";
import { InstanceUserData } from "./instanceUserData";
import { InstanceVolume } from "./instanceVolume";
import { IotDevice } from "./iotDevice";
import { IotHub } from "./iotHub";
import { IotNetwork } from "./iotNetwork";
import { IotRoute } from "./iotRoute";
import { K8sCluster } from "./k8sCluster";
import { K8sPool } from "./k8sPool";
import { Lb } from "./lb";
import { LbBackend } from "./lbBackend";
import { LbCertificate } from "./lbCertificate";
import { LbFrontend } from "./lbFrontend";
import { LbIp } from "./lbIp";
import { LbRoute } from "./lbRoute";
import { MnqCredential } from "./mnqCredential";
import { MnqNamespace } from "./mnqNamespace";
import { ObjectBucket } from "./objectBucket";
import { ObjectBucketAcl } from "./objectBucketAcl";
import { ObjectBucketLockConfiguration } from "./objectBucketLockConfiguration";
import { ObjectBucketPolicy } from "./objectBucketPolicy";
import { ObjectBucketWebsiteConfiguration } from "./objectBucketWebsiteConfiguration";
import { ObjectItem } from "./objectItem";
import { RdbAcl } from "./rdbAcl";
import { RdbDatabase } from "./rdbDatabase";
import { RdbDatabaseBackup } from "./rdbDatabaseBackup";
import { RdbInstance } from "./rdbInstance";
import { RdbPrivilege } from "./rdbPrivilege";
import { RdbReadReplica } from "./rdbReadReplica";
import { RdbUser } from "./rdbUser";
import { RedisCluster } from "./redisCluster";
import { RegistryNamespace } from "./registryNamespace";
import { Secret } from "./secret";
import { SecretVersion } from "./secretVersion";
import { TemDomain } from "./temDomain";
import { VpcGatewayNetwork } from "./vpcGatewayNetwork";
import { VpcPrivateNetwork } from "./vpcPrivateNetwork";
import { VpcPublicGateway } from "./vpcPublicGateway";
import { VpcPublicGatewayDhcp } from "./vpcPublicGatewayDhcp";
import { VpcPublicGatewayDhcpReservation } from "./vpcPublicGatewayDhcpReservation";
import { VpcPublicGatewayIp } from "./vpcPublicGatewayIp";
import { VpcPublicGatewayIpReverseDns } from "./vpcPublicGatewayIpReverseDns";
import { VpcPublicGatewayPatRule } from "./vpcPublicGatewayPatRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:index/accountProject:AccountProject":
                return new AccountProject(name, <any>undefined, { urn })
            case "scaleway:index/accountSshKey:AccountSshKey":
                return new AccountSshKey(name, <any>undefined, { urn })
            case "scaleway:index/appleSiliconServer:AppleSiliconServer":
                return new AppleSiliconServer(name, <any>undefined, { urn })
            case "scaleway:index/baremetalServer:BaremetalServer":
                return new BaremetalServer(name, <any>undefined, { urn })
            case "scaleway:index/cockpit:Cockpit":
                return new Cockpit(name, <any>undefined, { urn })
            case "scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser":
                return new CockpitGrafanaUser(name, <any>undefined, { urn })
            case "scaleway:index/cockpitToken:CockpitToken":
                return new CockpitToken(name, <any>undefined, { urn })
            case "scaleway:index/container:Container":
                return new Container(name, <any>undefined, { urn })
            case "scaleway:index/containerCron:ContainerCron":
                return new ContainerCron(name, <any>undefined, { urn })
            case "scaleway:index/containerDomain:ContainerDomain":
                return new ContainerDomain(name, <any>undefined, { urn })
            case "scaleway:index/containerNamespace:ContainerNamespace":
                return new ContainerNamespace(name, <any>undefined, { urn })
            case "scaleway:index/containerToken:ContainerToken":
                return new ContainerToken(name, <any>undefined, { urn })
            case "scaleway:index/domainRecord:DomainRecord":
                return new DomainRecord(name, <any>undefined, { urn })
            case "scaleway:index/domainZone:DomainZone":
                return new DomainZone(name, <any>undefined, { urn })
            case "scaleway:index/flexibleIp:FlexibleIp":
                return new FlexibleIp(name, <any>undefined, { urn })
            case "scaleway:index/function:Function":
                return new Function(name, <any>undefined, { urn })
            case "scaleway:index/functionCron:FunctionCron":
                return new FunctionCron(name, <any>undefined, { urn })
            case "scaleway:index/functionDomain:FunctionDomain":
                return new FunctionDomain(name, <any>undefined, { urn })
            case "scaleway:index/functionNamespace:FunctionNamespace":
                return new FunctionNamespace(name, <any>undefined, { urn })
            case "scaleway:index/functionToken:FunctionToken":
                return new FunctionToken(name, <any>undefined, { urn })
            case "scaleway:index/iamApiKey:IamApiKey":
                return new IamApiKey(name, <any>undefined, { urn })
            case "scaleway:index/iamApplication:IamApplication":
                return new IamApplication(name, <any>undefined, { urn })
            case "scaleway:index/iamGroup:IamGroup":
                return new IamGroup(name, <any>undefined, { urn })
            case "scaleway:index/iamPolicy:IamPolicy":
                return new IamPolicy(name, <any>undefined, { urn })
            case "scaleway:index/iamSshKey:IamSshKey":
                return new IamSshKey(name, <any>undefined, { urn })
            case "scaleway:index/instanceImage:InstanceImage":
                return new InstanceImage(name, <any>undefined, { urn })
            case "scaleway:index/instanceIp:InstanceIp":
                return new InstanceIp(name, <any>undefined, { urn })
            case "scaleway:index/instanceIpReverseDns:InstanceIpReverseDns":
                return new InstanceIpReverseDns(name, <any>undefined, { urn })
            case "scaleway:index/instancePlacementGroup:InstancePlacementGroup":
                return new InstancePlacementGroup(name, <any>undefined, { urn })
            case "scaleway:index/instancePrivateNic:InstancePrivateNic":
                return new InstancePrivateNic(name, <any>undefined, { urn })
            case "scaleway:index/instanceSecurityGroup:InstanceSecurityGroup":
                return new InstanceSecurityGroup(name, <any>undefined, { urn })
            case "scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules":
                return new InstanceSecurityGroupRules(name, <any>undefined, { urn })
            case "scaleway:index/instanceServer:InstanceServer":
                return new InstanceServer(name, <any>undefined, { urn })
            case "scaleway:index/instanceSnapshot:InstanceSnapshot":
                return new InstanceSnapshot(name, <any>undefined, { urn })
            case "scaleway:index/instanceUserData:InstanceUserData":
                return new InstanceUserData(name, <any>undefined, { urn })
            case "scaleway:index/instanceVolume:InstanceVolume":
                return new InstanceVolume(name, <any>undefined, { urn })
            case "scaleway:index/iotDevice:IotDevice":
                return new IotDevice(name, <any>undefined, { urn })
            case "scaleway:index/iotHub:IotHub":
                return new IotHub(name, <any>undefined, { urn })
            case "scaleway:index/iotNetwork:IotNetwork":
                return new IotNetwork(name, <any>undefined, { urn })
            case "scaleway:index/iotRoute:IotRoute":
                return new IotRoute(name, <any>undefined, { urn })
            case "scaleway:index/k8sCluster:K8sCluster":
                return new K8sCluster(name, <any>undefined, { urn })
            case "scaleway:index/k8sPool:K8sPool":
                return new K8sPool(name, <any>undefined, { urn })
            case "scaleway:index/lb:Lb":
                return new Lb(name, <any>undefined, { urn })
            case "scaleway:index/lbBackend:LbBackend":
                return new LbBackend(name, <any>undefined, { urn })
            case "scaleway:index/lbCertificate:LbCertificate":
                return new LbCertificate(name, <any>undefined, { urn })
            case "scaleway:index/lbFrontend:LbFrontend":
                return new LbFrontend(name, <any>undefined, { urn })
            case "scaleway:index/lbIp:LbIp":
                return new LbIp(name, <any>undefined, { urn })
            case "scaleway:index/lbRoute:LbRoute":
                return new LbRoute(name, <any>undefined, { urn })
            case "scaleway:index/mnqCredential:MnqCredential":
                return new MnqCredential(name, <any>undefined, { urn })
            case "scaleway:index/mnqNamespace:MnqNamespace":
                return new MnqNamespace(name, <any>undefined, { urn })
            case "scaleway:index/objectBucket:ObjectBucket":
                return new ObjectBucket(name, <any>undefined, { urn })
            case "scaleway:index/objectBucketAcl:ObjectBucketAcl":
                return new ObjectBucketAcl(name, <any>undefined, { urn })
            case "scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration":
                return new ObjectBucketLockConfiguration(name, <any>undefined, { urn })
            case "scaleway:index/objectBucketPolicy:ObjectBucketPolicy":
                return new ObjectBucketPolicy(name, <any>undefined, { urn })
            case "scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration":
                return new ObjectBucketWebsiteConfiguration(name, <any>undefined, { urn })
            case "scaleway:index/objectItem:ObjectItem":
                return new ObjectItem(name, <any>undefined, { urn })
            case "scaleway:index/rdbAcl:RdbAcl":
                return new RdbAcl(name, <any>undefined, { urn })
            case "scaleway:index/rdbDatabase:RdbDatabase":
                return new RdbDatabase(name, <any>undefined, { urn })
            case "scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup":
                return new RdbDatabaseBackup(name, <any>undefined, { urn })
            case "scaleway:index/rdbInstance:RdbInstance":
                return new RdbInstance(name, <any>undefined, { urn })
            case "scaleway:index/rdbPrivilege:RdbPrivilege":
                return new RdbPrivilege(name, <any>undefined, { urn })
            case "scaleway:index/rdbReadReplica:RdbReadReplica":
                return new RdbReadReplica(name, <any>undefined, { urn })
            case "scaleway:index/rdbUser:RdbUser":
                return new RdbUser(name, <any>undefined, { urn })
            case "scaleway:index/redisCluster:RedisCluster":
                return new RedisCluster(name, <any>undefined, { urn })
            case "scaleway:index/registryNamespace:RegistryNamespace":
                return new RegistryNamespace(name, <any>undefined, { urn })
            case "scaleway:index/secret:Secret":
                return new Secret(name, <any>undefined, { urn })
            case "scaleway:index/secretVersion:SecretVersion":
                return new SecretVersion(name, <any>undefined, { urn })
            case "scaleway:index/temDomain:TemDomain":
                return new TemDomain(name, <any>undefined, { urn })
            case "scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork":
                return new VpcGatewayNetwork(name, <any>undefined, { urn })
            case "scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork":
                return new VpcPrivateNetwork(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGateway:VpcPublicGateway":
                return new VpcPublicGateway(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp":
                return new VpcPublicGatewayDhcp(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation":
                return new VpcPublicGatewayDhcpReservation(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp":
                return new VpcPublicGatewayIp(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns":
                return new VpcPublicGatewayIpReverseDns(name, <any>undefined, { urn })
            case "scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule":
                return new VpcPublicGatewayPatRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "index/accountProject", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/accountSshKey", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/appleSiliconServer", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/baremetalServer", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/cockpit", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/cockpitGrafanaUser", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/cockpitToken", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/container", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/containerCron", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/containerDomain", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/containerNamespace", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/containerToken", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/domainRecord", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/domainZone", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/flexibleIp", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/function", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/functionCron", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/functionDomain", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/functionNamespace", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/functionToken", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iamApiKey", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iamApplication", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iamGroup", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iamPolicy", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iamSshKey", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceImage", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceIp", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceIpReverseDns", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instancePlacementGroup", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instancePrivateNic", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceSecurityGroup", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceSecurityGroupRules", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceServer", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceSnapshot", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceUserData", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/instanceVolume", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iotDevice", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iotHub", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iotNetwork", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/iotRoute", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/k8sCluster", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/k8sPool", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lb", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lbBackend", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lbCertificate", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lbFrontend", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lbIp", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/lbRoute", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/mnqCredential", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/mnqNamespace", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectBucket", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectBucketAcl", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectBucketLockConfiguration", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectBucketPolicy", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectBucketWebsiteConfiguration", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/objectItem", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbAcl", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbDatabase", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbDatabaseBackup", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbInstance", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbPrivilege", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbReadReplica", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/rdbUser", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/redisCluster", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/registryNamespace", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/secret", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/secretVersion", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/temDomain", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcGatewayNetwork", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPrivateNetwork", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGateway", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGatewayDhcp", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGatewayDhcpReservation", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGatewayIp", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGatewayIpReverseDns", _module)
pulumi.runtime.registerResourceModule("scaleway", "index/vpcPublicGatewayPatRule", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("scaleway", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:scaleway") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});