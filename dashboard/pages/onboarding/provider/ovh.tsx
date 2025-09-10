import OVHAccountDetails from '@components/account-details/OVHAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function CivoCredentials() {
  return (
    <ProviderContent
      provider={allProviders.OVH}
      providerName="OVH"
      description="OVH is an open and sustainability oriented cloud provider"
    >
      <OVHAccountDetails />
    </ProviderContent>
  );
}
