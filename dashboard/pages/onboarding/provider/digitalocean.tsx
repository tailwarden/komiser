import DigitalOceanAccountDetails from '@components/account-details/DigitalOceanAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function DigitalOceanCredentials() {
  return (
    <ProviderContent
      provider={allProviders.DIGITAL_OCEAN}
      providerName="DigitalOcean"
      description="DigitalOcean is a cloud hosting provider that offers cloud
  computing services and Infrastructure as a Service (IaaS)."
    >
      <DigitalOceanAccountDetails />
    </ProviderContent>
  );
}
