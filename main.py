import json
import pandas as pd
import plotly.express as px
import streamlit as st

FILENAME = "tmp/output.json"

class contestPerf:
    def __init__(self, name, time, rating):
        self.name, self.time, self.rating = name, time, rating
        

def main():
    print("Hello from profile-tracker!")
    with open(FILENAME,'r') as f:
        data =  json.load(f)
    perfs = []
    for dic in data["RatingChanges"]:
        perfs.append(contestPerf(dic["ContestName"], dic["RatingUpdateTimeSeconds"], dic["NewRating"]))
    perfs = sorted(perfs, key= lambda x : x.time)
    biodata = data["Biodata"]
    for k, v in biodata.items():
        if k != "TitlePhoto":
            st.text(f"{k} : {v}\n")
    st.image(
        biodata["TitlePhoto"], 
        caption="Profile", 
        width=400,
        alt=biodata["TitlePhoto"]
    )


    # 1. data
    df = pd.DataFrame(
        {"Times": [p.time for p in perfs], "ratings": [p.rating for p in perfs], "Names" : [p.name for p in perfs]}
    )

    # 2. Create Plotly line chart with text labels
    fig = px.line(
        df,
        x="Times",
        y="ratings",
        text="Names",
        markers=True,  # Shows dots at each data point
        title="CF performance",
    )

    # 3. Position and format the labels
    fig.update_traces(
        textposition="top center",  # Places labels above the points
        textfont=dict(size=12, color="black"),
    )

    # 4. Display in Streamlit
    st.plotly_chart(fig, use_container_width=True)


if __name__ == "__main__":
    main()
