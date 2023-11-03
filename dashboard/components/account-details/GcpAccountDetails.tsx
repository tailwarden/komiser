import { ChangeEvent, useRef, useState } from 'react';
import classNames from 'classnames';
import { GCPCredentials } from '@utils/cloudAccountHelpers';
import Folder2Icon from '../icons/Folder2Icon';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import InputFileSelect from '../onboarding-wizard/InputFileSelect';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface GcpAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<GCPCredentials>;
  hasError?: boolean;
}

function GcpAccountDetails({
  cloudAccountData,
  hasError = false
}: GcpAccountDetailsProps) {
  const [isValidationError, setIsValidationError] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [file, setFile] = useState<string>(
    cloudAccountData?.credentials.serviceAccountKeyPath || ''
  );

  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const handleButtonClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const fileName = event.target.files?.[0]?.name;

    if (fileName) {
      if (!fileName.endsWith('.json')) {
        setIsValidationError(true);
        setErrorMessage(
          'The chosen file is not supported. Please choose a different file for the credentials.'
        );
        setFile(fileName);
      }
    } else {
      setIsValidationError(true);
      setErrorMessage('Please choose a file.');
    }
  };

  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        label="Account name"
        placeholder="my-gcp-account"
        required
        value={cloudAccountData?.name}
      />

      <div
        className={classNames(
          'flex flex-col space-y-8 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <div>
          <InputFileSelect
            type="text"
            label="File path"
            id="file-path-input"
            name="serviceAccountKeyPath"
            icon={<Folder2Icon />}
            subLabel="Enter the path or browse the file"
            placeholder="C:\Documents\Komiser\credentials"
            fileInputRef={fileInputRef}
            iconClick={handleButtonClick}
            handleFileChange={handleFileChange}
            handleInputChange={e => setFile(e.target.value)}
            value={file}
            hasError={isValidationError}
            errorMessage={errorMessage}
          />
        </div>
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your GCP account. Please check if the file
          is correct.
        </div>
      )}
    </div>
  );
}

export default GcpAccountDetails;
