import Head from 'next/head';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';
import LabelledInput from '../../components/onboarding-wizard/LabelledInput';
import CredentialsButton from '../../components/onboarding-wizard/CredentialsButton';

export default function AzureCredentials() {
  const handleNext = () => {};

  return (
    <div>
      <Head>
        <title>Setup Civo - Komiser</title>
        <meta name="description" content="Setup Azure - Komiser" />
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
                icon={
                  <svg
                    width="24"
                    height="25"
                    viewBox="0 0 24 25"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5"
                  >
                    <path
                      d="M19.79 15.5621C17.73 17.6121 14.78 18.2421 12.19 17.4321L7.48002 22.1321C7.14002 22.4821 6.47002 22.6921 5.99002 22.6221L3.81002 22.3221C3.09002 22.2221 2.42002 21.5421 2.31002 20.8221L2.01002 18.6421C1.94002 18.1621 2.17002 17.4921 2.50002 17.1521L7.20002 12.4521C6.40002 9.85215 7.02002 6.90215 9.08002 4.85215C12.03 1.90215 16.82 1.90215 19.78 4.85215C22.74 7.80215 22.74 12.6121 19.79 15.5621Z"
                      stroke="#0C1717"
                      stroke-width="1.5"
                      stroke-miterlimit="10"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                    <path
                      d="M6.89001 18.1221L9.19001 20.4221"
                      stroke="#0C1717"
                      stroke-width="1.5"
                      stroke-miterlimit="10"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                    <path
                      d="M14.5 11.6318C15.3284 11.6318 16 10.9603 16 10.1318C16 9.30341 15.3284 8.63184 14.5 8.63184C13.6716 8.63184 13 9.30341 13 10.1318C13 10.9603 13.6716 11.6318 14.5 11.6318Z"
                      stroke="#0C1717"
                      stroke-width="1.5"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                  </svg>
                }
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
            <p>Civo Video Here</p>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
