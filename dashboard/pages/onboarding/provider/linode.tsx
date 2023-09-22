import LinodeAccountDetails from '@components/account-details/LinodeAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function LinodeCredentials() {
  return (
    <ProviderContent
      provider={allProviders.LINODE}
      providerName="Linode"
      description="Linode is a global cloud hosting provider offering
  infrastructure-as-a-service solutions. It provides virtual
  servers, storage, and related services for deploying and managing
  applications in the cloud."
    >
      <LinodeAccountDetails />
    </ProviderContent>
  );
}
