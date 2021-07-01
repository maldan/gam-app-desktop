export class Helper {
  static addEventListener(obj: any, events: string, fn: any) {
    events.split(' ').forEach((x) => {
      obj.addEventListener(x, fn);
    });
  }

  static removeEventListener(obj: any, events: string, fn: any) {
    events.split(' ').forEach((x) => {
      obj.removeEventListener(x, fn);
    });
  }

  static convertName(name: string) {
    return name
      .split('-')
      .slice(1)
      .filter((x) => x !== 'app' && x !== 'gam')
      .join(' ');
  }
}
