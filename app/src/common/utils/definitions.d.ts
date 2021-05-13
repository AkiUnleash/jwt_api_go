declare module "*.png"

declare module '*.jpg' {
  const value: any;
  export = value;
}

declare module '*.scss' {
  interface SubClass {
    [str: string]: string;
  }
  const classNames: SubClass;
  export = classNames;
}

declare module '*.css' {
  interface SubClass {
    [str: string]: string;
  }
  const classNames: SubClass;
  export = classNames;
}