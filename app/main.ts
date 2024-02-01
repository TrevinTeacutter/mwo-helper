import {app, BrowserWindow, ipcMain, screen} from 'electron';
import * as path from 'path';
import * as fs from 'fs';
import fetch from "node-fetch";

export interface Response {
  MatchDetails: MatchDetails,
  UserDetails: UserDetails[],
}

export interface MatchDetails {
  Map: string;
  ViewMode: string;
  TimeOfDay: string;
  GameMode: string;
  Region: string;
  MatchTimeMinutes: string;
  UseStockLoadout: boolean;
  NoMechQuirks: boolean;
  NoMechEfficiencies: boolean;
  WinningTeam: string;
  Team1Score: number;
  Team2Score: number;
  MatchDuration: string;
  CompleteTime: string;
}

export interface UserDetails {
  Username: string;
  IsSpectator: boolean;
  Team: string;
  Lance: string;
  MechItemID: number;
  MechName: string;
  SkillTier: number;
  HealthPercentage: number;
  Kills: number;
  KillsMostDamage: number;
  Assists: number;
  ComponentsDestroyed: number;
  MatchScore: number;
  Damage: number;
  TeamDamage: number;
  UnitTag: string;
}

let win: BrowserWindow | null = null;
const args = process.argv.slice(1),
  serve = args.some(val => val === '--serve');

function createWindow(): BrowserWindow {

  const size = screen.getPrimaryDisplay().workAreaSize;

  // Create the browser window.
  win = new BrowserWindow({
    x: 0,
    y: 0,
    width: size.width,
    height: size.height,
    webPreferences: {
      nodeIntegration: true,
      allowRunningInsecureContent: (serve),
      contextIsolation: false,
    },
  });

  if (serve) {
    const debug = require('electron-debug');
    debug();

    require('electron-reloader')(module);
    win.loadURL('http://localhost:4200');
  } else {
    // Path when running electron executable
    let pathIndex = './index.html';

    if (fs.existsSync(path.join(__dirname, '../dist/index.html'))) {
       // Path when running electron in local folder
      pathIndex = '../dist/index.html';
    }

    const url = new URL(path.join('file:', __dirname, pathIndex));
    win.loadURL(url.href);
  }

  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store window
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null;
  });

  return win;
}

try {
  // This method will be called when Electron has finished
  // initialization and is ready to create browser windows.
  // Some APIs can only be used after this event occurs.
  // Added 400 ms to fix the black background issue while using transparent window. More detais at https://github.com/electron/electron/issues/15947
  app.on('ready', () => setTimeout(createWindow, 400));

  // Quit when all windows are closed.
  app.on('window-all-closed', () => {
    // On OS X it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform !== 'darwin') {
      app.quit();
    }
  });

  app.on('activate', () => {
    // On OS X it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (win === null) {
      createWindow();
    }
  });

  app.whenReady().then(() => {
    ipcMain.handle("matchDetails", async (event, apiKey: string, matchID: string) => {
      console.debug(`apiKey(${apiKey}) matchID(${matchID})`)

      const response = await fetch(`https://mwomercs.com/api/v1/matches/${matchID}?api_token=${apiKey}`);
      const results: Response = JSON.parse(await response.text())

      return results;
    })
  })
} catch (e) {
  // Catch Error
  // throw e;
}

