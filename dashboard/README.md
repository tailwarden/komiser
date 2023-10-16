# 🚀 Komiser Dashboard

Komiser dashboard is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

**Full frontend stack:**

- 🖥 `Next.js`
- 📜 `Typescript`
- 🎨 `Tailwind`
- 📖 `Storybook`
- 🧪 `Jest`
- 📚 `React Testing Library`

## 🚀 Getting Started

First, run the development server:

```shell

# From the Komiser root folder start the Komiser server, run:

go run \*.go start --config /path/to/config.toml

# In a different terminal tab in the dashboard folder, run:

NEXT_PUBLIC_API_URL=http://localhost:3000 npm run dev

# Alternatively, you can create an .env file with it:

NEXT_PUBLIC_API_URL=http://localhost:3000
```

Open [http://localhost:3002/](http://localhost:3002). If you see the dashboard, 🎉 congrats! It's all up and running correctly.
![Dashboard Image](https://user-images.githubusercontent.com/13384559/224318056-3d2c68bc-aa56-49c8-841a-bb297e380dc9.png)

❗ If you get an error page such as this, please refer to the logs and our [docs](https://docs.komiser.io/docs/introduction/getting-started).
![Error Image](https://user-images.githubusercontent.com/13384559/224320642-0bf6814b-d97a-4ad9-95a0-ca82e353c5d0.png)

## 🧩 Components

Komiser components are documented under `/components`

> 💡 **Hint:**
> We have the following import aliases defined in `tsconfig.json`
> ```json
> {
> "@components/": "/dashboard/components/",
> "@services/": "/dashboard/services/",
> "@environments/": "/dashboard/environments/",
> "@utils/": "/dashboard/utils/",
> "@styles/": "/dashboard/styles/"
> }
> ```

**Component convention:**

- 📁 Component folder: component name in `kebab-case`
- 📄 Component file: component name in `UpperCamelCase.*`
- 📖 Component story: component name in `UpperCamelCase.stories.*`
- 🎭 Component story mock (if needed): component name in `UpperCamelCase.mocks.*`
- 🧪 Component unit test: component name in `UpperCamelCase.test.*`
- 🧐 Check `Card` example for more details:
  ![Component Example](https://user-images.githubusercontent.com/13384559/224307211-2ce62245-de24-4ee7-a156-fb54d8d34b4f.png)

**Additional instructions:**

- 📖 To view this component on Storybook, run: `npm run storybook`, then pick `Card`
  ![Storybook Image](https://user-images.githubusercontent.com/13384559/224320112-e21d2ed4-1e22-4a33-adb3-6c236c4d4208.png)
- 🧪 To run the unit tests, run: `npm run test:watch`, hit `p`, then `card`
  ![Unit Test Image](https://user-images.githubusercontent.com/13384559/224320260-19b1359e-1bfb-4db5-8379-918dacd7da44.png)

## 🧪 Testing

We use Jest & React Testing Library for our unit tests.

**Testing convention:**

- ✅ All new Utils need to be tested. Existing ones when being changed
- ✅ All tests should be wrapped in a `describe`
- ✅ If it's a unit test for a function: `describe('[replace with function name]', () => { ... })`
- ✅ If it's a unit test for a util: `describe('[replace with util name] util', () => { ... })`
- ✅ If it's a unit test for a component: `describe('[replace with component name]', () => { ... })`
- ✅ A test should use 'it' for the test function: `it('should do something', () => { ... })`

**Testing examples:**

- Simple Jest unit test example (snippet from `/utils/formatNumber.test.ts`):

```typescript
import formatNumber from './formatNumber';

describe('formatNumber util', () => {
  it('should format number (over a thousand) in short notation', () => {
    const result = formatNumber(12345);
    expect(result).toBe('12K');
  });
  ...
});
```

- Jest & Testing library example (snippet from `/components/card/Card.test.tsx`):

```typescript
import { render, screen } from '@testing-library/react';
import RefreshIcon from '../icons/RefreshIcon';
import Card from './Card';

describe('Card', () => {
  it('should render card component without crashing', () => {
    render(
      <Card
      label="Test card"
      value={500}
      icon={<RefreshIcon width={24} height={24} />}
      />
    );
  });

  it('should display the value formatted', () => {
    render(
      <Card
      label="Test card"
      value={5000}
      icon={<RefreshIcon width={24} height={24} />}
      />
    );
    const formattedNumber = screen.getByTestId('formattedNumber');
    expect(formattedNumber).toHaveTextContent('5K');
  });
  ...
});
```

If you're looking for an example with event firing and state updates, have a look at `components/select-checkbox/SelectCheckbox.test.tsx`:
```typescript
it('opens the dropdown when clicked', () => {
  const { getByRole, getByText } = render(
    <SelectCheckbox
    label="Test Label"
    query="provider"
    exclude={[]}
    setExclude={() => {}}
    />
  );

  fireEvent.click(getByRole('button'));

  expect(getByText('Item 1')).toBeInTheDocument();
  expect(getByText('Item 2')).toBeInTheDocument();
  expect(getByText('Item 3')).toBeInTheDocument();
});
```

## 🤝 Contributing

We welcome all contributors to join us on the mission of improving Komiser, especially when it comes to writing tests and adding documentation.

Not sure where to start?

- 📖 Read the [contributor guidelines](https://docs.komiser.io/docs/introduction/community)
- 💬 [Join our Discord](https://discord.tailwarden.com/) and hang with us on #contributors channel.

## 📚 Learn More

To learn more about our stack, take a look at the following resources:

- [Next.js documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.
- [Tailwind documentation](https://tailwindcss.com/docs/)
- [Storybook documentation](https://storybook.js.org/docs/react/get-started/whats-a-story)
- [Jest documentation](https://jestjs.io/docs/getting-started)
- [React testing library documentation](https://testing-library.com/docs/dom-testing-library/intro)

## 🎥 Walkthrough video

[![Watch the video](https://komiser-assets-cdn.s3.eu-central-1.amazonaws.com/images/dashboard-contrib-video-thumb.png)](https://www.youtube.com/watch?v=uwxj11-eRt8)
