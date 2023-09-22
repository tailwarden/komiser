import ScalewayAccountDetails from '@components/account-details/ScalewayAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function ScalewayCredentials() {
  return (
    <ProviderContent
      provider={allProviders.SCALE_WAY}
      providerName="Scaleway"
      description="Scaleway is a cloud infrastructure provider that offers a range of
  cloud resources, including bare metal servers, virtual private
  servers, object storage, and more"
    >
      <ScalewayAccountDetails />
    </ProviderContent>
  );
}
