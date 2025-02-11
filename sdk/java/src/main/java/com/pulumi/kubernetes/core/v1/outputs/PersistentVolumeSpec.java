// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.AWSElasticBlockStoreVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.AzureDiskVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.AzureFilePersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.CSIPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.CephFSPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.CinderPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.FCVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.FlexPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.FlockerVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.GCEPersistentDiskVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.GlusterfsPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.HostPathVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.ISCSIPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.LocalVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.NFSVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.ObjectReference;
import com.pulumi.kubernetes.core.v1.outputs.PhotonPersistentDiskVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.PortworxVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.QuobyteVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.RBDPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.ScaleIOPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.StorageOSPersistentVolumeSource;
import com.pulumi.kubernetes.core.v1.outputs.VolumeNodeAffinity;
import com.pulumi.kubernetes.core.v1.outputs.VsphereVirtualDiskVolumeSource;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PersistentVolumeSpec {
    /**
     * @return accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
     * 
     */
    private @Nullable List<String> accessModes;
    /**
     * @return awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     * 
     */
    private @Nullable AWSElasticBlockStoreVolumeSource awsElasticBlockStore;
    /**
     * @return azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
     * 
     */
    private @Nullable AzureDiskVolumeSource azureDisk;
    /**
     * @return azureFile represents an Azure File Service mount on the host and bind mount to the pod.
     * 
     */
    private @Nullable AzureFilePersistentVolumeSource azureFile;
    /**
     * @return capacity is the description of the persistent volume&#39;s resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
     * 
     */
    private @Nullable Map<String,String> capacity;
    /**
     * @return cephFS represents a Ceph FS mount on the host that shares a pod&#39;s lifetime
     * 
     */
    private @Nullable CephFSPersistentVolumeSource cephfs;
    /**
     * @return cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     * 
     */
    private @Nullable CinderPersistentVolumeSource cinder;
    /**
     * @return claimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
     * 
     */
    private @Nullable ObjectReference claimRef;
    /**
     * @return csi represents storage that is handled by an external CSI driver (Beta feature).
     * 
     */
    private @Nullable CSIPersistentVolumeSource csi;
    /**
     * @return fc represents a Fibre Channel resource that is attached to a kubelet&#39;s host machine and then exposed to the pod.
     * 
     */
    private @Nullable FCVolumeSource fc;
    /**
     * @return flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
     * 
     */
    private @Nullable FlexPersistentVolumeSource flexVolume;
    /**
     * @return flocker represents a Flocker volume attached to a kubelet&#39;s host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
     * 
     */
    private @Nullable FlockerVolumeSource flocker;
    /**
     * @return gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     * 
     */
    private @Nullable GCEPersistentDiskVolumeSource gcePersistentDisk;
    /**
     * @return glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
     * 
     */
    private @Nullable GlusterfsPersistentVolumeSource glusterfs;
    /**
     * @return hostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     * 
     */
    private @Nullable HostPathVolumeSource hostPath;
    /**
     * @return iscsi represents an ISCSI Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. Provisioned by an admin.
     * 
     */
    private @Nullable ISCSIPersistentVolumeSource iscsi;
    /**
     * @return local represents directly-attached storage with node affinity
     * 
     */
    private @Nullable LocalVolumeSource local;
    /**
     * @return mountOptions is the list of mount options, e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
     * 
     */
    private @Nullable List<String> mountOptions;
    /**
     * @return nfs represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     * 
     */
    private @Nullable NFSVolumeSource nfs;
    /**
     * @return nodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
     * 
     */
    private @Nullable VolumeNodeAffinity nodeAffinity;
    /**
     * @return persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
     * 
     */
    private @Nullable String persistentVolumeReclaimPolicy;
    /**
     * @return photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
     * 
     */
    private @Nullable PhotonPersistentDiskVolumeSource photonPersistentDisk;
    /**
     * @return portworxVolume represents a portworx volume attached and mounted on kubelets host machine
     * 
     */
    private @Nullable PortworxVolumeSource portworxVolume;
    /**
     * @return quobyte represents a Quobyte mount on the host that shares a pod&#39;s lifetime
     * 
     */
    private @Nullable QuobyteVolumeSource quobyte;
    /**
     * @return rbd represents a Rados Block Device mount on the host that shares a pod&#39;s lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
     * 
     */
    private @Nullable RBDPersistentVolumeSource rbd;
    /**
     * @return scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
     * 
     */
    private @Nullable ScaleIOPersistentVolumeSource scaleIO;
    /**
     * @return storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
     * 
     */
    private @Nullable String storageClassName;
    /**
     * @return storageOS represents a StorageOS volume that is attached to the kubelet&#39;s host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
     * 
     */
    private @Nullable StorageOSPersistentVolumeSource storageos;
    /**
     * @return volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
     * 
     */
    private @Nullable String volumeMode;
    /**
     * @return vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
     * 
     */
    private @Nullable VsphereVirtualDiskVolumeSource vsphereVolume;

    private PersistentVolumeSpec() {}
    /**
     * @return accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
     * 
     */
    public List<String> accessModes() {
        return this.accessModes == null ? List.of() : this.accessModes;
    }
    /**
     * @return awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     * 
     */
    public Optional<AWSElasticBlockStoreVolumeSource> awsElasticBlockStore() {
        return Optional.ofNullable(this.awsElasticBlockStore);
    }
    /**
     * @return azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
     * 
     */
    public Optional<AzureDiskVolumeSource> azureDisk() {
        return Optional.ofNullable(this.azureDisk);
    }
    /**
     * @return azureFile represents an Azure File Service mount on the host and bind mount to the pod.
     * 
     */
    public Optional<AzureFilePersistentVolumeSource> azureFile() {
        return Optional.ofNullable(this.azureFile);
    }
    /**
     * @return capacity is the description of the persistent volume&#39;s resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
     * 
     */
    public Map<String,String> capacity() {
        return this.capacity == null ? Map.of() : this.capacity;
    }
    /**
     * @return cephFS represents a Ceph FS mount on the host that shares a pod&#39;s lifetime
     * 
     */
    public Optional<CephFSPersistentVolumeSource> cephfs() {
        return Optional.ofNullable(this.cephfs);
    }
    /**
     * @return cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     * 
     */
    public Optional<CinderPersistentVolumeSource> cinder() {
        return Optional.ofNullable(this.cinder);
    }
    /**
     * @return claimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
     * 
     */
    public Optional<ObjectReference> claimRef() {
        return Optional.ofNullable(this.claimRef);
    }
    /**
     * @return csi represents storage that is handled by an external CSI driver (Beta feature).
     * 
     */
    public Optional<CSIPersistentVolumeSource> csi() {
        return Optional.ofNullable(this.csi);
    }
    /**
     * @return fc represents a Fibre Channel resource that is attached to a kubelet&#39;s host machine and then exposed to the pod.
     * 
     */
    public Optional<FCVolumeSource> fc() {
        return Optional.ofNullable(this.fc);
    }
    /**
     * @return flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
     * 
     */
    public Optional<FlexPersistentVolumeSource> flexVolume() {
        return Optional.ofNullable(this.flexVolume);
    }
    /**
     * @return flocker represents a Flocker volume attached to a kubelet&#39;s host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
     * 
     */
    public Optional<FlockerVolumeSource> flocker() {
        return Optional.ofNullable(this.flocker);
    }
    /**
     * @return gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     * 
     */
    public Optional<GCEPersistentDiskVolumeSource> gcePersistentDisk() {
        return Optional.ofNullable(this.gcePersistentDisk);
    }
    /**
     * @return glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
     * 
     */
    public Optional<GlusterfsPersistentVolumeSource> glusterfs() {
        return Optional.ofNullable(this.glusterfs);
    }
    /**
     * @return hostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     * 
     */
    public Optional<HostPathVolumeSource> hostPath() {
        return Optional.ofNullable(this.hostPath);
    }
    /**
     * @return iscsi represents an ISCSI Disk resource that is attached to a kubelet&#39;s host machine and then exposed to the pod. Provisioned by an admin.
     * 
     */
    public Optional<ISCSIPersistentVolumeSource> iscsi() {
        return Optional.ofNullable(this.iscsi);
    }
    /**
     * @return local represents directly-attached storage with node affinity
     * 
     */
    public Optional<LocalVolumeSource> local() {
        return Optional.ofNullable(this.local);
    }
    /**
     * @return mountOptions is the list of mount options, e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
     * 
     */
    public List<String> mountOptions() {
        return this.mountOptions == null ? List.of() : this.mountOptions;
    }
    /**
     * @return nfs represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     * 
     */
    public Optional<NFSVolumeSource> nfs() {
        return Optional.ofNullable(this.nfs);
    }
    /**
     * @return nodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
     * 
     */
    public Optional<VolumeNodeAffinity> nodeAffinity() {
        return Optional.ofNullable(this.nodeAffinity);
    }
    /**
     * @return persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
     * 
     */
    public Optional<String> persistentVolumeReclaimPolicy() {
        return Optional.ofNullable(this.persistentVolumeReclaimPolicy);
    }
    /**
     * @return photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
     * 
     */
    public Optional<PhotonPersistentDiskVolumeSource> photonPersistentDisk() {
        return Optional.ofNullable(this.photonPersistentDisk);
    }
    /**
     * @return portworxVolume represents a portworx volume attached and mounted on kubelets host machine
     * 
     */
    public Optional<PortworxVolumeSource> portworxVolume() {
        return Optional.ofNullable(this.portworxVolume);
    }
    /**
     * @return quobyte represents a Quobyte mount on the host that shares a pod&#39;s lifetime
     * 
     */
    public Optional<QuobyteVolumeSource> quobyte() {
        return Optional.ofNullable(this.quobyte);
    }
    /**
     * @return rbd represents a Rados Block Device mount on the host that shares a pod&#39;s lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
     * 
     */
    public Optional<RBDPersistentVolumeSource> rbd() {
        return Optional.ofNullable(this.rbd);
    }
    /**
     * @return scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
     * 
     */
    public Optional<ScaleIOPersistentVolumeSource> scaleIO() {
        return Optional.ofNullable(this.scaleIO);
    }
    /**
     * @return storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
     * 
     */
    public Optional<String> storageClassName() {
        return Optional.ofNullable(this.storageClassName);
    }
    /**
     * @return storageOS represents a StorageOS volume that is attached to the kubelet&#39;s host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
     * 
     */
    public Optional<StorageOSPersistentVolumeSource> storageos() {
        return Optional.ofNullable(this.storageos);
    }
    /**
     * @return volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
     * 
     */
    public Optional<String> volumeMode() {
        return Optional.ofNullable(this.volumeMode);
    }
    /**
     * @return vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
     * 
     */
    public Optional<VsphereVirtualDiskVolumeSource> vsphereVolume() {
        return Optional.ofNullable(this.vsphereVolume);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PersistentVolumeSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> accessModes;
        private @Nullable AWSElasticBlockStoreVolumeSource awsElasticBlockStore;
        private @Nullable AzureDiskVolumeSource azureDisk;
        private @Nullable AzureFilePersistentVolumeSource azureFile;
        private @Nullable Map<String,String> capacity;
        private @Nullable CephFSPersistentVolumeSource cephfs;
        private @Nullable CinderPersistentVolumeSource cinder;
        private @Nullable ObjectReference claimRef;
        private @Nullable CSIPersistentVolumeSource csi;
        private @Nullable FCVolumeSource fc;
        private @Nullable FlexPersistentVolumeSource flexVolume;
        private @Nullable FlockerVolumeSource flocker;
        private @Nullable GCEPersistentDiskVolumeSource gcePersistentDisk;
        private @Nullable GlusterfsPersistentVolumeSource glusterfs;
        private @Nullable HostPathVolumeSource hostPath;
        private @Nullable ISCSIPersistentVolumeSource iscsi;
        private @Nullable LocalVolumeSource local;
        private @Nullable List<String> mountOptions;
        private @Nullable NFSVolumeSource nfs;
        private @Nullable VolumeNodeAffinity nodeAffinity;
        private @Nullable String persistentVolumeReclaimPolicy;
        private @Nullable PhotonPersistentDiskVolumeSource photonPersistentDisk;
        private @Nullable PortworxVolumeSource portworxVolume;
        private @Nullable QuobyteVolumeSource quobyte;
        private @Nullable RBDPersistentVolumeSource rbd;
        private @Nullable ScaleIOPersistentVolumeSource scaleIO;
        private @Nullable String storageClassName;
        private @Nullable StorageOSPersistentVolumeSource storageos;
        private @Nullable String volumeMode;
        private @Nullable VsphereVirtualDiskVolumeSource vsphereVolume;
        public Builder() {}
        public Builder(PersistentVolumeSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessModes = defaults.accessModes;
    	      this.awsElasticBlockStore = defaults.awsElasticBlockStore;
    	      this.azureDisk = defaults.azureDisk;
    	      this.azureFile = defaults.azureFile;
    	      this.capacity = defaults.capacity;
    	      this.cephfs = defaults.cephfs;
    	      this.cinder = defaults.cinder;
    	      this.claimRef = defaults.claimRef;
    	      this.csi = defaults.csi;
    	      this.fc = defaults.fc;
    	      this.flexVolume = defaults.flexVolume;
    	      this.flocker = defaults.flocker;
    	      this.gcePersistentDisk = defaults.gcePersistentDisk;
    	      this.glusterfs = defaults.glusterfs;
    	      this.hostPath = defaults.hostPath;
    	      this.iscsi = defaults.iscsi;
    	      this.local = defaults.local;
    	      this.mountOptions = defaults.mountOptions;
    	      this.nfs = defaults.nfs;
    	      this.nodeAffinity = defaults.nodeAffinity;
    	      this.persistentVolumeReclaimPolicy = defaults.persistentVolumeReclaimPolicy;
    	      this.photonPersistentDisk = defaults.photonPersistentDisk;
    	      this.portworxVolume = defaults.portworxVolume;
    	      this.quobyte = defaults.quobyte;
    	      this.rbd = defaults.rbd;
    	      this.scaleIO = defaults.scaleIO;
    	      this.storageClassName = defaults.storageClassName;
    	      this.storageos = defaults.storageos;
    	      this.volumeMode = defaults.volumeMode;
    	      this.vsphereVolume = defaults.vsphereVolume;
        }

        @CustomType.Setter
        public Builder accessModes(@Nullable List<String> accessModes) {
            this.accessModes = accessModes;
            return this;
        }
        public Builder accessModes(String... accessModes) {
            return accessModes(List.of(accessModes));
        }
        @CustomType.Setter
        public Builder awsElasticBlockStore(@Nullable AWSElasticBlockStoreVolumeSource awsElasticBlockStore) {
            this.awsElasticBlockStore = awsElasticBlockStore;
            return this;
        }
        @CustomType.Setter
        public Builder azureDisk(@Nullable AzureDiskVolumeSource azureDisk) {
            this.azureDisk = azureDisk;
            return this;
        }
        @CustomType.Setter
        public Builder azureFile(@Nullable AzureFilePersistentVolumeSource azureFile) {
            this.azureFile = azureFile;
            return this;
        }
        @CustomType.Setter
        public Builder capacity(@Nullable Map<String,String> capacity) {
            this.capacity = capacity;
            return this;
        }
        @CustomType.Setter
        public Builder cephfs(@Nullable CephFSPersistentVolumeSource cephfs) {
            this.cephfs = cephfs;
            return this;
        }
        @CustomType.Setter
        public Builder cinder(@Nullable CinderPersistentVolumeSource cinder) {
            this.cinder = cinder;
            return this;
        }
        @CustomType.Setter
        public Builder claimRef(@Nullable ObjectReference claimRef) {
            this.claimRef = claimRef;
            return this;
        }
        @CustomType.Setter
        public Builder csi(@Nullable CSIPersistentVolumeSource csi) {
            this.csi = csi;
            return this;
        }
        @CustomType.Setter
        public Builder fc(@Nullable FCVolumeSource fc) {
            this.fc = fc;
            return this;
        }
        @CustomType.Setter
        public Builder flexVolume(@Nullable FlexPersistentVolumeSource flexVolume) {
            this.flexVolume = flexVolume;
            return this;
        }
        @CustomType.Setter
        public Builder flocker(@Nullable FlockerVolumeSource flocker) {
            this.flocker = flocker;
            return this;
        }
        @CustomType.Setter
        public Builder gcePersistentDisk(@Nullable GCEPersistentDiskVolumeSource gcePersistentDisk) {
            this.gcePersistentDisk = gcePersistentDisk;
            return this;
        }
        @CustomType.Setter
        public Builder glusterfs(@Nullable GlusterfsPersistentVolumeSource glusterfs) {
            this.glusterfs = glusterfs;
            return this;
        }
        @CustomType.Setter
        public Builder hostPath(@Nullable HostPathVolumeSource hostPath) {
            this.hostPath = hostPath;
            return this;
        }
        @CustomType.Setter
        public Builder iscsi(@Nullable ISCSIPersistentVolumeSource iscsi) {
            this.iscsi = iscsi;
            return this;
        }
        @CustomType.Setter
        public Builder local(@Nullable LocalVolumeSource local) {
            this.local = local;
            return this;
        }
        @CustomType.Setter
        public Builder mountOptions(@Nullable List<String> mountOptions) {
            this.mountOptions = mountOptions;
            return this;
        }
        public Builder mountOptions(String... mountOptions) {
            return mountOptions(List.of(mountOptions));
        }
        @CustomType.Setter
        public Builder nfs(@Nullable NFSVolumeSource nfs) {
            this.nfs = nfs;
            return this;
        }
        @CustomType.Setter
        public Builder nodeAffinity(@Nullable VolumeNodeAffinity nodeAffinity) {
            this.nodeAffinity = nodeAffinity;
            return this;
        }
        @CustomType.Setter
        public Builder persistentVolumeReclaimPolicy(@Nullable String persistentVolumeReclaimPolicy) {
            this.persistentVolumeReclaimPolicy = persistentVolumeReclaimPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder photonPersistentDisk(@Nullable PhotonPersistentDiskVolumeSource photonPersistentDisk) {
            this.photonPersistentDisk = photonPersistentDisk;
            return this;
        }
        @CustomType.Setter
        public Builder portworxVolume(@Nullable PortworxVolumeSource portworxVolume) {
            this.portworxVolume = portworxVolume;
            return this;
        }
        @CustomType.Setter
        public Builder quobyte(@Nullable QuobyteVolumeSource quobyte) {
            this.quobyte = quobyte;
            return this;
        }
        @CustomType.Setter
        public Builder rbd(@Nullable RBDPersistentVolumeSource rbd) {
            this.rbd = rbd;
            return this;
        }
        @CustomType.Setter
        public Builder scaleIO(@Nullable ScaleIOPersistentVolumeSource scaleIO) {
            this.scaleIO = scaleIO;
            return this;
        }
        @CustomType.Setter
        public Builder storageClassName(@Nullable String storageClassName) {
            this.storageClassName = storageClassName;
            return this;
        }
        @CustomType.Setter
        public Builder storageos(@Nullable StorageOSPersistentVolumeSource storageos) {
            this.storageos = storageos;
            return this;
        }
        @CustomType.Setter
        public Builder volumeMode(@Nullable String volumeMode) {
            this.volumeMode = volumeMode;
            return this;
        }
        @CustomType.Setter
        public Builder vsphereVolume(@Nullable VsphereVirtualDiskVolumeSource vsphereVolume) {
            this.vsphereVolume = vsphereVolume;
            return this;
        }
        public PersistentVolumeSpec build() {
            final var o = new PersistentVolumeSpec();
            o.accessModes = accessModes;
            o.awsElasticBlockStore = awsElasticBlockStore;
            o.azureDisk = azureDisk;
            o.azureFile = azureFile;
            o.capacity = capacity;
            o.cephfs = cephfs;
            o.cinder = cinder;
            o.claimRef = claimRef;
            o.csi = csi;
            o.fc = fc;
            o.flexVolume = flexVolume;
            o.flocker = flocker;
            o.gcePersistentDisk = gcePersistentDisk;
            o.glusterfs = glusterfs;
            o.hostPath = hostPath;
            o.iscsi = iscsi;
            o.local = local;
            o.mountOptions = mountOptions;
            o.nfs = nfs;
            o.nodeAffinity = nodeAffinity;
            o.persistentVolumeReclaimPolicy = persistentVolumeReclaimPolicy;
            o.photonPersistentDisk = photonPersistentDisk;
            o.portworxVolume = portworxVolume;
            o.quobyte = quobyte;
            o.rbd = rbd;
            o.scaleIO = scaleIO;
            o.storageClassName = storageClassName;
            o.storageos = storageos;
            o.volumeMode = volumeMode;
            o.vsphereVolume = vsphereVolume;
            return o;
        }
    }
}
