import HetznerAccountDetails from '@components/account-details/HetznerAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function HetznerCredentials() {
  return (
    <ProviderContent
      provider={allProviders.HETZNER}
      providerName="Hetzner"
      description="Hetzner is a cloud hosting provider that offers cloud servers, volumes, load balancers and networking via Hetzner Cloud API."
    >
      <HetznerAccountDetails />
    </ProviderContent>
  );
}
