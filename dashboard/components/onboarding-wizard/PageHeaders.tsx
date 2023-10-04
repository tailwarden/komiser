import Image from 'next/image';

function OnboardingWizardHeader({ title }: { title: string }) {
  return (
    <>
      <Image
        src="/assets/img/komiser-logo.svg"
        alt="Komiser Logo"
        width={0}
        height={0}
        className="h-20 w-40"
      />
      <div className="mt-3 py-2 text-2xl font-semibold leading-8">{title}</div>
    </>
  );
}

export default OnboardingWizardHeader;
