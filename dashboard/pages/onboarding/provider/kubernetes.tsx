import Head from 'next/head';
import { useRouter } from 'next/router';
import { ReactNode, useRef, useState } from 'react';

import { allProviders } from '../../../utils/providerHelper';

import Folder2Icon from '../../../components/icons/Folder2Icon';
import DocumentTextIcon from '../../../components/icons/DocumentTextIcon';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import SelectInput from '../../../components/onboarding-wizard/SelectInput';
import PurplinCloud from '../../../components/onboarding-wizard/PurplinCloud';
import LabelledInput from '../../../components/onboarding-wizard/LabelledInput';
import InputFileSelect from '../../../components/onboarding-wizard/InputFileSelect';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';

interface SelectOptions {
  icon: ReactNode;
  label: string;
  value: string;
}

const options: SelectOptions[] = [
  {
    icon: <DocumentTextIcon />,
    label: 'Credentials File',
    value: 'credentials-file'
  }
];

export default function KubernetesCredentials() {
  const provider = allProviders.KUBERNETES;

  const router = useRouter();
  const [credentialType, setCredentialType] = useState<string>(
    options[0].value
  );

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
    router.push(`/onboarding/${provider}`);
  };

  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const handleButtonClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: any) => {
    const file = event.target.files[0];
    // TODO: (onboarding-wizard) handle file change and naming. Set Input field to file.name and use temporary file path for the upload value
    console.log(file);
  };

  function handleSelectChange(newValue: string) {
    setCredentialType(newValue);
  }

  return (
    <div>
      <Head>
        <title>Setup Kubernetes - Komiser</title>
        <meta name="description" content="Setup Kubernetes - Komiser" />
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
              <SelectInput
                icon="Change"
                label={'Source'}
                displayValues={options}
                value={credentialType}
                handleChange={handleSelectChange}
                values={options.map(option => option.value)}
              />
              <InputFileSelect
                type="text"
                id="file-path-input"
                label="File path"
                subLabel="Enter the path or browse the file"
                placeholder="C:\Documents\Komiser\credentials"
                icon={<Folder2Icon className="h-6 w-6" />}
                fileInputRef={fileInputRef}
                iconClick={handleButtonClick}
                handleFileChange={handleFileChange}
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
