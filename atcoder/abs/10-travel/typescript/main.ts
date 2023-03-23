// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// atcoder problem slolver template
import * as fs from "fs"

/**
 * stdIO : parameter read and solve outputs
 *
 *
 */
class stdIO {
  // property
  // for stdin
  inputs: string = ""
  inputArray: string[]
  inputIndex: number = 0;

  // for stdout
  outputBuffer: string = "";

  // constructor]
  // readline from stdin and split this
  constructor() {
    this.inputs = (fs.readFileSync(process.stdin.fd, "utf8"))
    this.inputArray = this.inputs.split(/\s/)
    this.inputIndex = 0
  }

  // input meethod
  getNext(): string {
    return this.inputArray[this.inputIndex++];
  }

  getNextNum(): number {
    return Number(this.getNext());
  }

  getNums(N: number): number[] {
    const arr: number[] = [];
    for (let i: number = 0; i < N; ++i) {
      arr[i] = this.getNextNum();
    }
    return arr;
  }

  // bufferd output methods
  toString(...args: any[]): string {
    const buff: string[] = []
    args.forEach((arg: any) => {
      buff.push(this._toString(arg))
    })
    const buff2 = buff.join(' ')
    return buff2;
  }

  // 出力
  print(...args: any): void {
    this.outputBuffer += this.toString(...args)
  }

  p(...args: any[]): void {
    this.print(...args)
  }

  println(...args: any): void {
    this.print(...args);
    this.print("\n")
  }


  flush(): void {
    console.log(this.outputBuffer.trim())
    this.outputBuffer = ""
  }

  // 引数を文字列か
  private _toString(arg1: any, quote: boolean = false): string {
    let buff = ""
    if (arg1 === null) {
      buff = "null"
    } else if (arg1 == "") {
    } else if (arg1 == "\n") {
      buff = "\n"
    } else if (Array.isArray(arg1)) {
      buff = this._arrayToString(arg1)
    } else {
      switch (typeof arg1) {
        case "undefined":
          buff = (typeof arg1);
          break
        case "string":
          if (quote) {
            buff = "'" + arg1 + "'";
          } else {
            buff = arg1
          }
          break
        default:
          buff = arg1.toString()
          break
      }
    }
    return buff;
  }

  // 配列を文字列に
  private _arrayToString(arr: any[]): string {
    const buff: string[] = []
    arr.map((e) => {
      buff.push(this._toString(e, true))
    })
    const buff2 = "[ " + buff.join(", ") + " ]"
    return buff2;
  }
}

// definition
// stop points
class TPoint {
  t: number;
  x: number;
  y: number;

  constructor(t: number = 0, x: number = 0, y: number = 0) {
    this.t = t
    this.x = x
    this.y = y
  }

  toString() {
    return "[" + this.t + ", " + this.x + ", " + this.y + "]";
  }

  tdistance(tp: TPoint): number {
    const td: number = Math.abs(this.t - tp.t)
    return td
  }

  distance(tp: TPoint): number {
    const xd: number = Math.abs(this.x - tp.x)
    const yd: number = Math.abs(this.y - tp.y)
    return xd + yd
  }
}

//
class TPList {
  points: TPoint[];

  constructor() {
    this.points = []
    this.points.push(new TPoint())
  }

  add(t: number, x: number, y: number) {
    const tp = new TPoint(t, x, y)
    this.points.push(tp)
  }

  readN(n: number) {
    for (let i = 1; i <= n; i++) {
      const t = io.getNextNum()
      const x = io.getNextNum()
      const y = io.getNextNum()

      this.add(t, x, y)
    }
  }

  getN(n: number): TPoint {
    return this.points[n]
  }

  toString() {
    let buff: string = ""
    for (let tp of this.points) {
      buff += tp.toString() + "\n"
    }
    return buff.trim()
  }
}

// solver functions
function canMoveTP(tp: TPList, N: number): boolean {
  for (let i = 1; i <= N; i++) {
    const tp0 = tp.getN(i - 1)
    const tp1 = tp.getN(i)
    const d = tp0.tdistance(tp1) - tp0.distance(tp1)
    // io.println(tp0 + " " + tp1 + " " + d + "\n")
    if ((d < 0) || (d % 2 > 0)) {
      return false
    }
  }
  return true
}



// main routin
const io: stdIO = new stdIO()
const TP0 = new TPoint

solve()
io.flush()

// solve main()
function solve() {
  // initialize
  const tp = new TPList()

  // input
  const N: number = io.getNextNum()
  tp.readN(N)

  // solve
  const canMove: boolean = canMoveTP(tp, N)

  // output
  io.print(canMove ? "Yes" : "No")
}

