export type TemplateProps = {
  sampleTextProp: string;
};

function Template({ sampleTextProp }: TemplateProps) {
  return (
    <div className="rounded-lg bg-primary p-6 text-sm text-white">
      {sampleTextProp}
    </div>
  );
}

export default Template;
