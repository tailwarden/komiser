import KubernetesAccountDetails from '@components/account-details/KubernetesAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function KubernetesCredentials() {
  return (
    <ProviderContent
      provider={allProviders.KUBERNETES}
      providerName="Kubernetes"
      description="Kubernetes, also known as K8s, is an open-source system for
  automating deployment, scaling, and management of containerized
  applications."
    >
      <KubernetesAccountDetails />
    </ProviderContent>
  );
}
