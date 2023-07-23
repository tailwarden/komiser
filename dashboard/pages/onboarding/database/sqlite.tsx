import Head from 'next/head';
import { useRef } from 'react';

import { allDBProviders } from '../../../utils/providerHelper';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import Folder2Icon from '../../../components/icons/Folder2Icon';
import DatabasePurplin from '../../../components/onboarding-wizard/DatabasePurplin';
import InputFileSelect from '../../../components/onboarding-wizard/InputFileSelect';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';

export default function SqliteCredentials() {
  const database = allDBProviders.SQLITE;

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

  return (
    <div>
      <Head>
        <title>Configure SQLite - Komiser</title>
        <meta name="description" content="Setup SQLite - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your SQLite database"
          progressBarWidth="81%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              SQLite is a lightweight, serverless, self-contained RDBMS that
              operates directly on files. It is known for its simplicity, ease
              of use, and portability across platforms.
            </div>
          </div>

          <div className="flex flex-col space-y-4 py-10">
            <div className="space-y-[0.2]">
              <InputFileSelect
                type="text"
                id="file-path-input"
                label="File path"
                subLabel="Enter the path or browse the file"
                placeholder="C:\Documents\Komiser\database"
                icon={<Folder2Icon className="h-6 w-6" />}
                fileInputRef={fileInputRef}
                iconClick={handleButtonClick}
                handleFileChange={handleFileChange}
              />
            </div>
          </div>

          <CredentialsButton handleNext={handleNext} nextLabel="Add database" />
        </LeftSideLayout>

        <RightSideLayout>
          <DatabasePurplin database={database} />
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
