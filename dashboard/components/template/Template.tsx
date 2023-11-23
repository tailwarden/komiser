export type TemplateProps = {
  sampleTextProp: string;
};

function Template({ sampleTextProp }: TemplateProps) {
  return (
    <div className="rounded-lg bg-darkcyan-500 p-6 text-sm text-white">
      {sampleTextProp}
    </div>
  );
}

export default Template;
