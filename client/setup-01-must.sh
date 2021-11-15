# typescript

yarn add -D typescript @types/node @types/react @types/react-dom

# tsconfig.json
cat <<EOS | jq > tsconfig.json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    },
    "target": "es2020",
    "skipLibCheck": true,
    "strict": true,
    "noEmit": true,
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "jsx": "preserve",
    "lib": ["dom", "dom.iterable", "esnext"],
    "allowJs": true,
    "forceConsistentCasingInFileNames": true
  },
  "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx"],
  "exclude": ["node_modules"]
}

EOS

# rm unused file
rm -rf pages/ styles/
mkdir -p src/pages

# src/pages/index.tsx
cat << EOS > src/pages/index.tsx
import { NextPage } from 'next'

const Index: NextPage = () => {
    return <div>hoge</div>
}

export default Index;
EOS

# install jest
yarn add jest @types/jest ts-jest -D

# jest.config

cat << EOS | jq >  jest.config.js
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  globals: {
    'ts-jest': {
      diagnostics: false,
    },
  },
}
EOS

# .eslintrc.json

# eslint
yarn add -D eslint @typescript-eslint/parser @typescript-eslint/eslint-plugin eslint-config-prettier eslint-plugin-jsx-a11y eslint-plugin-prettier eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-jest

cat << EOS | jq > .eslintrc.json
{
  "root": true,
  "plugins": ["jest", "react"],
  "env": {
    "es2020": true,
    "jest/globals": true
  },
  "extends": [
    "eslint:recommended",
    "plugin:prettier/recommended",
    "plugin:@typescript-eslint/eslint-recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended"
  ],
  "rules": {
    "no-mixed-operators": "error",
    "no-console": "off",
    "no-undef": "off",
    "react/jsx-uses-vars": 1,
    "@typescript-eslint/explicit-function-return-type": "off",
    "@typescript-eslint/explicit-module-boundary-types": "off",
    "@typescript-eslint/no-non-null-assertion": "off"
  },
  "parserOptions": {
    "project": "./tsconfig.json"
  },
  "parser": "@typescript-eslint/parser"
}
EOS

#

cat << EOS > .eslintignore
jest.config.js

EOS


# .prettierrc

yarn add -D prettier pretty-quick

cat << EOS | jq > .prettierrc
{
  "tabWidth": 2,
  "singleQuote": true,
  "semi": false,
  "trailingComma": "all",
  "printWidth": 100,
  "useTabs": false
}

EOS

# .prettierignore 作成するだけ
touch .prettierignore

# src/pages/_app.tsx
cat << EOS >  src/pages/_app.tsx
import React from 'react'
import { AppProps } from 'next/app'
import Head from 'next/head'

function App({ Component, pageProps }: AppProps): JSX.Element {
  React.useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    if (jssStyles) {
      jssStyles.parentElement?.removeChild(jssStyles)
    }
  }, [])

  return (
    <React.Fragment>
      <Head>
        <title>hayashiki | scaffold </title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <Component {...pageProps} />
    </React.Fragment>
  )
}

export default App

EOS

# huskey
yarn add --dev husky lint-staged

# form関連
yarn add formik yup
yarn add -D @types/yup


# 以下をたすか？
# extends": ["plugin:jest/recommended"]
