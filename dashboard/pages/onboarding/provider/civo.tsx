import Head from 'next/head';

import { allProviders } from '../../../utils/providerHelper';

import RecordCircleIcon from '../../../components/icons/RecordCircleIcon';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import PurplinCloud from '../../../components/onboarding-wizard/PurplinCloud';
import LabelledInput from '../../../components/onboarding-wizard/LabelledInput';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';

export default function CivoCredentials() {
  const provider = allProviders.CIVO;

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
  };

  return (
    <div>
      <Head>
        <title>Setup Civo - Komiser</title>
        <meta name="description" content="Setup Civo - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your Civo account"
          progressBarWidth="45%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              Civo is the first cloud native service provider powered only by
              Kubernetes.
            </div>
            <div>
              Read our guide on{' '}
              <a
                target="_blank"
                href="https://docs.komiser.io/docs/cloud-providers/civo"
                className="text-komiser-600"
                rel="noreferrer"
              >
                adding a CIVO account to Komiser.
              </a>
            </div>
          </div>

          <div className="flex flex-col space-y-4 py-10">
            <LabelledInput
              type="text"
              id="account-name"
              label="Account name"
              placeholder="my-civo-account"
            />

            <div className="flex flex-col space-y-[0.2] rounded-md bg-komiser-100 p-5">
              <LabelledInput
                type="text"
                id="source"
                label="Source"
                value="API Token"
                disabled={true}
                icon={<RecordCircleIcon />}
              />
              <LabelledInput
                type="text"
                id="api-token"
                label="API token"
                subLabel="An API key that is unique to your account"
                placeholder="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
              />
            </div>
          </div>

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
