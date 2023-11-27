import React, { useState } from 'react';
import Head from 'next/head';

import { Credentials, configureAccount } from '@utils/cloudAccountHelpers';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '@components/onboarding-wizard/OnboardingWizardLayout';
import PurplinCloud from '@components/onboarding-wizard/PurplinCloud';
import CredentialsButton from '@components/onboarding-wizard/CredentialsButton';
import Toast from '@components/toast/Toast';
import { Provider } from '@utils/providerHelper';
import { CloudAccountPayload } from '@components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import { useToast } from '@components/toast/ToastProvider';

interface ChildProps {
  cloudAccountData?: CloudAccountPayload<Credentials>;
  hasError?: boolean;
}

interface ProviderContentProps {
  provider: Provider;
  providerName: string;
  description: string;
  children: React.ReactElement<ChildProps>;
}

export default function ProviderContent({
  provider,
  providerName,
  description,
  children
}: ProviderContentProps) {
  const { toast, showToast, dismissToast } = useToast();

  const [hasError, setHasError] = useState(false);

  return (
    <div>
      <Head>
        <title>Setup {providerName} - Komiser</title>
        <meta name="description" content={`Setup ${providerName} - Komiser`} />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title={`Configure your ${providerName} account`}
          progressBarWidth="45%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">{description}</div>
            <div>
              Read our guide on{' '}
              <a
                target="_blank"
                href={`https://docs.komiser.io/configuration/cloud-providers/${provider}`}
                className="text-darkcyan-500"
                rel="noreferrer"
              >
                adding an {providerName} account to Komiser.
              </a>
            </div>
          </div>
          <form
            onSubmit={event =>
              configureAccount(event, provider, showToast, setHasError)
            }
          >
            {React.isValidElement(children)
              ? React.cloneElement(children, { hasError })
              : children}
            <CredentialsButton />
          </form>
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
