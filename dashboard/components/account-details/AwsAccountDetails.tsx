import { ChangeEvent, ReactNode, useRef, useState } from 'react';
import Folder2Icon from '../icons/Folder2Icon';
import SelectInput from '../onboarding-wizard/SelectInput';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import InputFileSelect from '../onboarding-wizard/InputFileSelect';
import KeyIcon from '../icons/KeyIcon';
import VariableIcon from '../icons/VariableIcon';
import DocumentTextIcon from '../icons/DocumentTextIcon';
import ShieldSecurityIcon from '../icons/ShieldSecurityIcon';
import { CloudAccount } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface SelectOptions {
  icon: ReactNode;
  label: string;
  value: string;
}

interface AwsAccountDetailsProps {
  cloudAccountData: CloudAccount;
  setCloudAccountData: (formData: CloudAccount) => void;
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

function AwsAccountDetails({
  cloudAccountData,
  setCloudAccountData
}: AwsAccountDetailsProps) {
  const [credentialType, setCredentialType] = useState<string>(
    options.find(option => option.value === cloudAccountData.credentials.source)
      ?.value ?? options[0].value
  );

  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const handleButtonClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  function handleNameChange(event: ChangeEvent<HTMLInputElement>) {
    setCloudAccountData({
      ...cloudAccountData,
      name: event?.target.value
    });
  }

  function handleSelectChange(newValue: string) {
    setCredentialType(newValue);

    setCloudAccountData({
      ...cloudAccountData,
      credentials: {
        ...cloudAccountData.credentials,
        source: newValue
      }
    });
  }

  function handleProfileChange(event: ChangeEvent<HTMLInputElement>) {
    setCloudAccountData({
      ...cloudAccountData,
      credentials: {
        ...cloudAccountData.credentials,
        profile: event?.target.value
      }
    });
  }

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];

    if (!file) return;

    setCloudAccountData({
      ...cloudAccountData,
      credentials: {
        ...cloudAccountData.credentials,
        path: file.name
      }
    });
  };

  return (
    <div className="flex flex-col space-y-8 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        label="Account name"
        placeholder="my-aws-account"
        value={cloudAccountData.name}
        onChange={handleNameChange}
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
          {[options[2].value, options[3].value].includes(credentialType) && (
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
              value={cloudAccountData.credentials.path}
            />
            <LabelledInput
              type="text"
              id="profile"
              label="Profile"
              placeholder="default"
              subLabel="Name of the section in the credentials file"
              value={cloudAccountData.credentials.profile}
              onChange={handleProfileChange}
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
  );
}

export default AwsAccountDetails;
