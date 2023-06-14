import { useRouter } from 'next/router';

import Button from '../button/Button';

function CredentialsButton({ handleNext }: { handleNext?: (e?: any) => void }) {
  const router = useRouter();

  return (
    <div className="flex justify-between">
      <Button
        onClick={() => router.back()}
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
  );
}

export default CredentialsButton;
