import Head from 'next/head';
import { useRouter } from 'next/router';
import { useState } from 'react';

import { allProviders } from '../../../utils/providerHelper';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import PurplinCloud from '../../../components/onboarding-wizard/PurplinCloud';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';
import AwsAccountDetails from '../../../components/account-details/AwsAccountDetails';
import { CloudAccount } from '../../../components/cloud-account/hooks/useCloudAccounts/useCloudAccount';

export default function AWSCredentials() {
  const provider = allProviders.AWS;

  const [cloudAccountData, setCloudAccountData] = useState<CloudAccount>({
    credentials: {
      path: '',
      profile: '',
      source: ''
    },
    name: '',
    provider: 'aws'
  });

  const router = useRouter();

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
  };

  return (
    <div>
      <Head>
        <title>Setup AWS - Komiser</title>
        <meta name="description" content="Setup AWS - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your AWS account"
          progressBarWidth="45%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              AWS is a cloud computing platform that provides infrastructure
              services, application services, and developer tools provided by
              Amazon.
            </div>
            <div>
              Read our guide on{' '}
              <a
                target="_blank"
                href="https://docs.komiser.io/docs/cloud-providers/aws"
                className="text-komiser-600"
                rel="noreferrer"
              >
                adding an AWS account to Komiser.
              </a>
            </div>
          </div>
          <AwsAccountDetails
            cloudAccountData={cloudAccountData}
            setCloudAccountData={setCloudAccountData}
          />
          <CredentialsButton handleNext={handleNext} />
        </LeftSideLayout>

        <RightSideLayout>
          <div className="relative">
            <PurplinCloud provider={provider} />
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
