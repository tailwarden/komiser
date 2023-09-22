import AzureAccountDetails from '@components/account-details/AzureAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function AzureCredentials() {
  return (
    <ProviderContent
      provider={allProviders.AZURE}
      providerName="Microsoft Azure"
      description="Microsoft Azure is Microsoft's public cloud computing
  platform. It provides a broad range of cloud services, including
  compute, analytics, storage and networking. Users can pick and
  choose from these services to develop and scale new applications
  or run existing applications in the public cloud."
    >
      <AzureAccountDetails />
    </ProviderContent>
  );
}
