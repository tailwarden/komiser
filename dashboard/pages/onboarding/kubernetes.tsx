import { useState } from 'react';
import Head from 'next/head';
import { useRouter } from 'next/router';

import Button from '../../components/button/Button';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';
import LabelledInput from '../../components/onboarding-wizard/LabelledInput';

export default function AWSCredentials() {
  const router = useRouter();
  const [provider, setProvider] = useState('aws');

  const handleNext = () => {
    router.push(`/onboarding/${provider}`);
  };

  const handleSuggest = () =>
    router.replace(
      'https://docs.komiser.io/docs/faqs#how-can-i-request-a-new-feature'
    );

  return (
    <div>
      <Head>
        <title>Setup Kubernetes - Komiser</title>
        <meta name="description" content="Setup AWS - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your Kubernetes account"
          progressBarWidth="45%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              Kubernetes, also known as K8s, is an open-source system for
              automating deployment, scaling, and management of containerized
              applications.
            </div>
            <div>
              Read our guide on{' '}
              <a
                target="_blank"
                href="https://docs.komiser.io/docs/cloud-providers/kubernetes"
                className="text-komiser-600"
                rel="noreferrer"
              >
                adding a Kubernetes account to Komiser.
              </a>
            </div>
          </div>
          <div className="flex flex-col space-y-8 py-10">
            <LabelledInput
              type="text"
              id="account-name"
              label="Account name"
              placeholder="my-kubernetes-account"
            />
            <div className="flex flex-col space-y-8 bg-komiser-100 p-4">
              <div>
                <label htmlFor="input-group-1" className="mb-2 block">
                  Source
                </label>
                <div className="relative mb-6">
                  <div className="pointer-events-none absolute inset-y-0 left-1 flex items-center pl-3">
                    <svg
                      width="24"
                      height="25"
                      viewBox="0 0 24 25"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5 text-gray-400"
                    >
                      <path
                        d="M21 7.63184V17.6318C21 20.6318 19.5 22.6318 16 22.6318H8C4.5 22.6318 3 20.6318 3 17.6318V7.63184C3 4.63184 4.5 2.63184 8 2.63184H16C19.5 2.63184 21 4.63184 21 7.63184Z"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-miterlimit="10"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M14.5 5.13184V7.13184C14.5 8.23184 15.4 9.13184 16.5 9.13184H18.5"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-miterlimit="10"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M8 13.6318H12"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-miterlimit="10"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M8 17.6318H16"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-miterlimit="10"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                    </svg>
                  </div>

                  <input
                    type="text"
                    id="input-group-1"
                    className="block w-full rounded border py-4 pl-12 text-sm text-black-900 outline outline-black-200 focus:outline-2 focus:outline-primary"
                    placeholder="Credentials File"
                  />
                  <button className="absolute inset-y-0 right-5 flex items-center pl-3 text-komiser-600">
                    Change
                  </button>
                </div>
              </div>
              <div>
                <input type="file" name="select-file" id="credential-file" />
              </div>
            </div>
          </div>
          <div className="flex justify-between">
            <Button
              onClick={handleSuggest}
              size="lg"
              style="text"
              type="button"
            >
              Back
            </Button>
            <Button
              onClick={handleNext}
              size="lg"
              style="primary"
              type="button"
              disabled={true}
            >
              Add a cloud account
            </Button>
          </div>
        </LeftSideLayout>

        <RightSideLayout>
          <div className="relative">
            <p>Kubernetes Video Here</p>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
