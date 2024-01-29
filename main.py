import json
import sys
import tkinter
import tkinter.font as tkfont

import requests


class Scoreboard(tkinter.Frame):
    HEADERS = ['Pilot', 'Lance', 'Mech', 'Kills', 'Assists', 'Damage', 'Team Damage', 'Status', 'Health', 'Match Score']

    def __init__(self, parent):
        tkinter.Frame.__init__(self, parent)

        self.configure(highlightbackground="grey", highlightthickness=2, padx=2, pady=2)
        self.header_frame = tkinter.Frame(self, highlightbackground="black", highlightthickness=2, padx=2, pady=2)
        self.data_frame = tkinter.Frame(self, highlightbackground="black", highlightthickness=2, padx=2, pady=2)

        self.header_frame.pack()
        self.data_frame.pack()

        for index, header in enumerate(Scoreboard.HEADERS):
            self.headers_entry = tkinter.Entry(self.header_frame, width=20, fg='White', font=('Arial', 16, 'bold'))
            self.headers_entry.grid(row=0, column=index)
            self.headers_entry.insert(tkinter.END, header)

        self.data_entry = None
        self.rows = []

        self.adjust()

    def update_rows(self, rows):
        self.rows = rows

        self.adjust()

    def adjust(self):
        self.data_frame.children.clear()

        for i, row in enumerate(self.rows):
            for j in range(len(Scoreboard.HEADERS)):
                self.data_entry = tkinter.Entry(self.data_frame, width=20, fg='White', font=('Arial', 16, 'bold'))
                self.data_entry.grid(row=i, column=j)
                self.data_entry.insert(tkinter.END, row[j])


class MatchOverview(tkinter.Frame):
    HEADERS = ['Map', 'Game Mode', 'Region', 'Time Limit', 'Winner', 'Team 1 Score', 'Team 2 Score', 'Duration']

    def __init__(self, parent):
        tkinter.Frame.__init__(self, parent)

        self.configure(highlightbackground="grey", highlightthickness=2, padx=2, pady=2)
        self.header_frame = tkinter.Frame(self, highlightbackground="black", highlightthickness=2, padx=2, pady=2)
        self.data_frame = tkinter.Frame(self, highlightbackground="black", highlightthickness=2, padx=2, pady=2)

        self.header_frame.pack()
        self.data_frame.pack()

        for index, header in enumerate(MatchOverview.HEADERS):
            self.headers_entry = tkinter.Entry(self.header_frame, width=20, fg='White', font=('Arial', 16, 'bold'))
            self.headers_entry.grid(row=0, column=index)
            self.headers_entry.insert(tkinter.END, header)

        self.data_entry = None
        self.data = []

        self.adjust()

    def update_data(self, data):
        self.data = data

        self.adjust()

    def adjust(self):
        self.data_frame.children.clear()

        for index, data in enumerate(self.data):
            self.data_entry = tkinter.Entry(self.data_frame, width=20, fg='White', font=('Arial', 16, 'bold'))
            self.data_entry.grid(row=0, column=index)
            self.data_entry.insert(tkinter.END, data)


def main(*args) -> int:
    window = tkinter.Tk()
    window.title("MWO Helper")
    window.eval('tk::PlaceWindow . center')

    font = tkfont.Font(family='Arial', size=16, weight='bold')

    header = tkinter.Frame()
    match_id_label = tkinter.Label(header, text="Match ID:", font=font)
    match_id_input = tkinter.Text(header, height=1, font=font)
    api_key_label = tkinter.Label(header, text="API Key:", font=font)
    api_key_input = tkinter.Text(header, height=1, font=font)

    match_overview = MatchOverview(window)
    team_1_scoreboard = Scoreboard(window)
    team_2_scoreboard = Scoreboard(window)

    def handle_submit_button():
        # matchID = '58877118108527'
        # api_key = 'Wc9zbe6Ti3dRSxEwyqexcEtIKtKJowdZCmMP2WUjbAmSRrVcSkd1xrzGUt0q'
        match_id = match_id_input.get(1.0, "end-1c")
        api_key = api_key_input.get(1.0, "end-1c")
        response = requests.get(f'https://mwomercs.com/api/v1/matches/{match_id}?api_token={api_key}')
        results = json.loads(response.text)

        if not results['MatchDetails'] is None:
            match_details = results['MatchDetails']
            match_overview.update_data([
                match_details["Map"],
                match_details["GameMode"],
                match_details["Region"],
                f'''{match_details["MatchTimeMinutes"]} minutes''',
                match_details["WinningTeam"],
                match_details["Team1Score"],
                match_details["Team2Score"],
                f'''{match_details["MatchDuration"]} seconds''',
            ])

        team1 = []
        team2 = []

        for userDetail in results['UserDetails']:
            if not userDetail['IsSpectator'] is False:
                continue

            details = [
                userDetail['Username'],
                userDetail['Lance'],
                str.upper(userDetail['MechName']),
                userDetail['Kills'],
                userDetail['Assists'],
                userDetail['Damage'],
                userDetail['TeamDamage'],
                'Alive' if userDetail['HealthPercentage'] > 0 else 'Dead',
                f'''{userDetail['HealthPercentage']}%''',
                userDetail['MatchScore'],
            ]

            if userDetail['Team'] == '1':
                team1.append(details)
                continue

            if userDetail['Team'] == '2':
                team2.append(details)
                continue

        team_1_scoreboard.update_rows(team1)
        team_2_scoreboard.update_rows(team2)

    button = tkinter.Button(header, text='Submit', command=handle_submit_button)

    api_key_label.grid(row=0, column=0)
    api_key_input.grid(row=0, column=1)
    match_id_label.grid(row=0, column=2)
    match_id_input.grid(row=0, column=3)
    button.grid(row=0, column=4)

    header.grid(row=0, column=0)
    match_overview.grid(row=1, column=0)
    team_1_scoreboard.grid(row=2, column=0)
    team_2_scoreboard.grid(row=3, column=0)

    window.mainloop()

    return 0


if __name__ == '__main__':
    sys.exit(main(sys.argv[1:]))

# {
#   "MatchDetails": {
#     "Map": "ForestColony",
#     "ViewMode": "Both",
#     "TimeOfDay": "Random",
#     "GameMode": "Assault",
#     "Region": "NorthAmerica",
#     "MatchTimeMinutes": "15",
#     "UseStockLoadout": false,
#     "NoMechQuirks": false,
#     "NoMechEfficiencies": false,
#     "WinningTeam": "2",
#     "Team1Score": 0,
#     "Team2Score": 1,
#     "MatchDuration": "69",
#     "CompleteTime": "2017-07-14T18:09:54+00:00"
#   },
#   "UserDetails": [
#     {
#       "Username": "pgtest901",
#       "IsSpectator": false,
#       "Team": "2",
#       "Lance": "1",
#       "MechItemID": 532,
#       "MechName": "hbr-fc",
#       "SkillTier": 5,
#       "HealthPercentage": 100,
#       "Kills": 0,
#       "KillsMostDamage": 0,
#       "Assists": 0,
#       "ComponentsDestroyed": 0,
#       "MatchScore": 28,
#       "Damage": 0,
#       "TeamDamage": 0,
#       "UnitTag": "_2ez"
#     },
#     {
#       "Username": "pgtest903",
#       "IsSpectator": false,
#       "Team": "1",
#       "Lance": "1",
#       "MechItemID": 355,
#       "MechName": "dwf-wc",
#       "SkillTier": 5,
#       "HealthPercentage": 0,
#       "Kills": 0,
#       "KillsMostDamage": 0,
#       "Assists": 0,
#       "ComponentsDestroyed": 0,
#       "MatchScore": 0,
#       "Damage": 0,
#       "TeamDamage": 0,
#       "UnitTag": ""
#     }
#   ]
# }
