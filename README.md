# quake-parser
Parser for Quake game log


## Usage
```bash
go test
go run .

```



## Report Example

```json
{
  "Report": [
    {
      "GameId": 18,
      "Total_kills": 489,
      "Players": [
        "Mal",
        "Zeh",
        "Dono da Bola",
        "Isgalamido",
        "Assasinu Credi",
        "Oootsimo"
      ],
      "Kills": [
        {
          "Name": "Mal",
          "Kills": 2
        },
        {
          "Name": "Zeh",
          "Kills": 20
        },
        {
          "Name": "Isgalamido",
          "Kills": 14
        },
        {
          "Name": "Oootsimo",
          "Kills": 10
        },
        {
          "Name": "Dono da Bola",
          "Kills": 14
        },
        {
          "Name": "Assasinu Credi",
          "Kills": 9
        }
      ],
      "Kill_by_means": [
        {
          "Name": "MOD_ROCKET",
          "Kills": 27
        },
        {
          "Name": "MOD_ROCKET_SPLASH",
          "Kills": 32
        },
        {
          "Name": "MOD_SHOTGUN",
          "Kills": 6
        },
        {
          "Name": "MOD_RAILGUN",
          "Kills": 10
        },
        {
          "Name": "MOD_MACHINEGUN",
          "Kills": 7
        },
        {
          "Name": "MOD_FALLING",
          "Kills": 1
        },
        {
          "Name": "MOD_TRIGGER_HURT",
          "Kills": 12
        }
      ]
    },
    {
      "GameId": 19,
      "Total_kills": 492,
      "Players": [
        "Dono da Bola",
        "Zeh",
        "Oootsimo",
        "Assasinu Credi"
      ],
      "Kills": [
        {
          "Name": "Dono da Bola",
          "Kills": 2
        },
        {
          "Name": "Oootsimo",
          "Kills": 1
        }
      ],
      "Kill_by_means": [
        {
          "Name": "MOD_ROCKET_SPLASH",
          "Kills": 2
        },
        {
          "Name": "MOD_ROCKET",
          "Kills": 1
        }
      ]
    },
    {
      "GameId": 20,
      "Total_kills": 589,
      "Players": [
        "Dono da Bola",
        "Isgalamido",
        "Zeh",
        "Oootsimo",
        "Mal",
        "Assasinu Credi"
      ],
      "Kills": [
        {
          "Name": "Assasinu Credi",
          "Kills": 19
        },
        {
          "Name": "Oootsimo",
          "Kills": 22
        },
        {
          "Name": "Dono da Bola",
          "Kills": 14
        },
        {
          "Name": "Zeh",
          "Kills": 19
        },
        {
          "Name": "Mal",
          "Kills": 6
        },
        {
          "Name": "Isgalamido",
          "Kills": 17
        }
      ],
      "Kill_by_means": [
        {
          "Name": "MOD_TRIGGER_HURT",
          "Kills": 14
        },
        {
          "Name": "MOD_RAILGUN",
          "Kills": 9
        },
        {
          "Name": "MOD_ROCKET_SPLASH",
          "Kills": 60
        },
        {
          "Name": "MOD_MACHINEGUN",
          "Kills": 4
        },
        {
          "Name": "MOD_SHOTGUN",
          "Kills": 4
        },
        {
          "Name": "MOD_FALLING",
          "Kills": 3
        },
        {
          "Name": "MOD_ROCKET",
          "Kills": 37
        }
      ]
    }
  ]
}

```

