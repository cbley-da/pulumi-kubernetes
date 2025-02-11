// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.LifecycleHandler;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class Lifecycle {
    /**
     * @return PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     * 
     */
    private @Nullable LifecycleHandler postStart;
    /**
     * @return PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod&#39;s termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod&#39;s termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     * 
     */
    private @Nullable LifecycleHandler preStop;

    private Lifecycle() {}
    /**
     * @return PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     * 
     */
    public Optional<LifecycleHandler> postStart() {
        return Optional.ofNullable(this.postStart);
    }
    /**
     * @return PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod&#39;s termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod&#39;s termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     * 
     */
    public Optional<LifecycleHandler> preStop() {
        return Optional.ofNullable(this.preStop);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Lifecycle defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable LifecycleHandler postStart;
        private @Nullable LifecycleHandler preStop;
        public Builder() {}
        public Builder(Lifecycle defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.postStart = defaults.postStart;
    	      this.preStop = defaults.preStop;
        }

        @CustomType.Setter
        public Builder postStart(@Nullable LifecycleHandler postStart) {
            this.postStart = postStart;
            return this;
        }
        @CustomType.Setter
        public Builder preStop(@Nullable LifecycleHandler preStop) {
            this.preStop = preStop;
            return this;
        }
        public Lifecycle build() {
            final var o = new Lifecycle();
            o.postStart = postStart;
            o.preStop = preStop;
            return o;
        }
    }
}
