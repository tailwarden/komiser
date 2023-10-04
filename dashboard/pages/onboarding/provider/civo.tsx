import CivoAccountDetails from '@components/account-details/CivoAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function CivoCredentials() {
  return (
    <ProviderContent
      provider={allProviders.CIVO}
      providerName="Civo"
      description="Civo is the first cloud native service provider powered only by
  Kubernetes."
    >
      <CivoAccountDetails />
    </ProviderContent>
  );
}
