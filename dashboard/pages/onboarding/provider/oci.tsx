import { allProviders } from '@utils/providerHelper';
import OciAccountDetails from '@components/account-details/OciAccountDetails';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function OciCredentials() {
  return (
    <ProviderContent
      provider={allProviders.OCI}
      providerName="OCI"
      description=""
    >
      <OciAccountDetails />
    </ProviderContent>
  );
}
