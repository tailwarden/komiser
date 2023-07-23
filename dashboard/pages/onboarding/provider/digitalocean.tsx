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

export default function DigitalOceanCredentials() {
  const provider = allProviders.DIGITAL_OCEAN;

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
  };

  return (
    <div>
      <Head>
        <title>Setup DigitalOcean - Komiser</title>
        <meta name="description" content="Setup DigitalOcean - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your DigitalOcean account"
          progressBarWidth="45%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              DigitalOcean is a cloud hosting provider that offers cloud
              computing services and Infrastructure as a Service (IaaS).
            </div>
            <div>
              Read our guide on{' '}
              <a
                target="_blank"
                href="https://docs.komiser.io/docs/cloud-providers/digital-ocean"
                className="text-komiser-600"
                rel="noreferrer"
              >
                adding a DigitalOcean account to Komiser.
              </a>
            </div>
          </div>

          <div className="flex flex-col space-y-4 py-10">
            <LabelledInput
              type="text"
              id="account-name"
              label="Account name"
              placeholder="my-digitalocean-account"
            />

            <div className="flex flex-col space-y-[0.2] rounded-md bg-komiser-100 p-5">
              <LabelledInput
                type="text"
                id="source"
                label="Source"
                value="Personal Access Token"
                disabled={true}
                icon={<RecordCircleIcon />}
              />
              <LabelledInput
                type="text"
                id="personal-access-token"
                label="Personal access token"
                subLabel="Personal access tokens function like ordinary OAuth access tokens"
                placeholder="abcd1234efgh5678ijklmnop90qrstuv"
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
