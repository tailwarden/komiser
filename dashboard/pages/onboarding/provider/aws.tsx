import Head from 'next/head';
import { useRouter } from 'next/router';
import { ReactNode, useRef, useState } from 'react';

import { allProviders } from '../../../utils/providerHelper';

import KeyIcon from '../../../components/icons/KeyIcon';
import Folder2Icon from '../../../components/icons/Folder2Icon';
import VariableIcon from '../../../components/icons/VariableIcon';
import DocumentTextIcon from '../../../components/icons/DocumentTextIcon';
import ShieldSecurityIcon from '../../../components/icons/ShieldSecurityIcon';

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
  },
  {
    icon: <KeyIcon />,
    label: 'Credentials keys',
    value: 'credentials-keys'
  },
  {
    icon: <VariableIcon />,
    label: 'Environment Variables',
    value: 'environment-variables'
  },
  {
    icon: <ShieldSecurityIcon />,
    label: 'IAM Instance Role',
    value: 'iam-instance-role'
  }
];

export default function AWSCredentials() {
  const provider = allProviders.AWS;

  const router = useRouter();
  const [credentialType, setCredentialType] = useState<string>(
    options[0].value
  );

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
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
  };

  function handleSelectChange(newValue: string) {
    setCredentialType(newValue);
  }

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
          <div className="flex flex-col space-y-8 py-10">
            <LabelledInput
              type="text"
              id="account-name"
              label="Account name"
              placeholder="my-aws-account"
            />
            <div className="flex flex-col space-y-8 rounded-md bg-komiser-100 p-5">
              <div>
                <SelectInput
                  icon="Change"
                  label={'Source'}
                  displayValues={options}
                  value={credentialType}
                  handleChange={handleSelectChange}
                  values={options.map(option => option.value)}
                />
                {[options[2].value, options[3].value].includes(
                  credentialType
                ) && (
                  <div className="mt-2 text-sm text-black-400">
                    {credentialType === options[3].value
                      ? 'Komiser will fetch the credentials from AWS'
                      : 'Komiser will load credentials from AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY.'}
                  </div>
                )}
              </div>

              {credentialType === options[0].value && (
                <div>
                  <InputFileSelect
                    type="text"
                    label="File path"
                    id="file-path-input"
                    icon={<Folder2Icon />}
                    subLabel="Enter the path or browse the file"
                    placeholder="C:\Documents\Komiser\credentials"
                    fileInputRef={fileInputRef}
                    iconClick={handleButtonClick}
                    handleFileChange={handleFileChange}
                  />

                  <LabelledInput
                    type="text"
                    id="profile"
                    label="Profile"
                    placeholder="default"
                    subLabel="Name of the section in the credentials file"
                  />
                </div>
              )}

              {credentialType === options[1].value && (
                <div>
                  <LabelledInput
                    type="text"
                    id="access-key-id"
                    label="Access key ID"
                    placeholder="AKIABCDEFGHIJKLMN12"
                    subLabel="Unique identifier used to access AWS services"
                  />

                  <LabelledInput
                    type="text"
                    id="secret-access-key"
                    label="Secret access key"
                    placeholder="AbCdEfGhIjKlMnOpQrStUvWxYz0123456789AbCd"
                    subLabel="The secret access key is generated by AWS when an access key is created"
                  />
                </div>
              )}
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
