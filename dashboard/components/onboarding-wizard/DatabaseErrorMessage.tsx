function DatabaseErrorMessage() {
  return (
    <div className="mt-8 rounded-md bg-red-50 px-3 py-1.5 text-sm text-red-500">
      We&apos;re sorry, but we were unable to connect to your database using the
      information provided. Please ensure that the information are correct and
      try again. If you continue to experience issues, please{' '}
      <a
        href="https://discord.tailwarden.com"
        target="_blank"
        rel="noreferrer"
        className="underline"
      >
        contact our support
      </a>{' '}
      team for assistance.
    </div>
  );
}

export default DatabaseErrorMessage;
