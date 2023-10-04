import GcpAccountDetails from '@components/account-details/GcpAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function GcpCredentials() {
  return (
    <ProviderContent
      provider={allProviders.GCP}
      providerName="GCP"
      description="GCP is a cloud computing platform that provides infrastructure
  services, application services, and developer tools provided by
  Google."
    >
      <GcpAccountDetails />
    </ProviderContent>
  );
}
