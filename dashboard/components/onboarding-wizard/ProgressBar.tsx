function OnboardingWizardProgressBar({ width }: { width: string }) {
  return (
    <div className="h-1 w-full bg-gray-100">
      <div className="h-1 bg-komiser-600" style={{ width }}></div>
    </div>
  );
}

export default OnboardingWizardProgressBar;
