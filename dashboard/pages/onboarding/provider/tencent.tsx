import TencentAccountDetails from '@components/account-details/TencentAccountDetails';
import { allProviders } from '../../../utils/providerHelper';
import ProviderContent from '../../../components/onboarding-wizard/ProviderContent';

export default function TencentCredentials() {
  return (
    <ProviderContent
      provider={allProviders.TENCENT}
      providerName="Tencent"
      description="Tencent Cloud is China's leading public cloud service
  provider (CSP). Tencent Cloud is a secure, reliable and
  high-performance public CSP that integrates Tencent's
  infrastructure building capabilities with the advantages of its
  massive user platform and ecosystem."
    >
      <TencentAccountDetails />
    </ProviderContent>
  );
}
