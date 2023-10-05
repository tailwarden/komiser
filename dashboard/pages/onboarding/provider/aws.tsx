import { allProviders } from '@utils/providerHelper';
import AwsAccountDetails from '@components/account-details/AwsAccountDetails';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function AWSCredentials() {
  return (
    <ProviderContent
      provider={allProviders.AWS}
      providerName="AWS"
      description="AWS is a cloud computing platform that provides infrastructure
    services, application services, and developer tools provided by
    Amazon."
    >
      <AwsAccountDetails />
    </ProviderContent>
  );
}
