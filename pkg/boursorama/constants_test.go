package boursorama

const authHTML = `
	<ul class="password-input">
		<li data-matrix-list-item data-matrix-list-item-index="0">
			<button type="button" data-matrix-key="FIVE" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMS42IDM2LjFjLjMuNC43LjcgMS40LjcuOSAwIDEuNC0uNiAxLjQtMS41di01aC45djVjMCAxLjYtMSAyLjMtMi4zIDIuMy0uOCAwLTEuNC0uMi0xLjktLjh6Ii8+PHBhdGggZD0ibTIwLjcgMzQuMy0uNy44djIuNGgtLjl2LTcuMmguOXYzLjdsMy4yLTMuN2gxLjFsLTMgMy40IDMuMiAzLjhoLTEuMXoiLz48cGF0aCBkPSJtMjcuNyAzMC4zaC45djYuNGgzLjR2LjhoLTQuMnYtNy4yeiIvPjwvZz48cGF0aCBkPSJtMTcuNCAyMC4xYzEuMSAxLjYgMi42IDIuNSA0LjggMi41IDIuNSAwIDQuMy0xLjggNC4zLTQuMiAwLTIuNi0xLjgtNC4yLTQuMy00LjItMS42IDAtMi45LjUtNC4yIDEuN2wtMS0uNnYtOWgxMHYxLjNoLTguNXY2LjhjLjktLjggMi4zLTEuNiA0LjEtMS42IDIuOSAwIDUuNSAxLjkgNS41IDUuNSAwIDMuNC0yLjYgNS42LTUuOCA1LjYtMi45IDAtNC42LTEuMS01LjgtMi44eiIvPjwvZz48L3N2Zz4=">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="1">
			<button type="button" data-matrix-key="TWO" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMy45IDM1LjloLTMuNmwtLjYgMS42aC0xbDIuOS03LjJoMS4xbDIuOSA3LjJoLTF6bS0zLjMtLjhoM2wtMS41LTMuOXoiLz48cGF0aCBkPSJtMTguNyAzMC4zaDMuMmMxLjIgMCAyIC44IDIgMS44IDAgLjktLjYgMS41LTEuMyAxLjYuOC4xIDEuNC45IDEuNCAxLjggMCAxLjItLjggMS45LTIuMSAxLjloLTMuM3YtNy4xem0zIDMuMWMuOCAwIDEuMi0uNSAxLjItMS4yIDAtLjYtLjQtMS4yLTEuMi0xLjJoLTIuMnYyLjNoMi4yem0wIDMuM2MuOCAwIDEuMy0uNSAxLjMtMS4ycy0uNS0xLjItMS4zLTEuMmgtMi4ydjIuNWgyLjJ6Ii8+PHBhdGggZD0ibTI3LjMgMzMuOWMwLTIuMiAxLjYtMy43IDMuNy0zLjcgMS4zIDAgMi4yLjYgMi43IDEuNGwtLjguNGMtLjQtLjYtMS4yLTEtMi0xLTEuNiAwLTIuOCAxLjItMi44IDIuOXMxLjIgMi45IDIuOCAyLjljLjggMCAxLjYtLjQgMi0xbC44LjRjLS42LjgtMS41IDEuNC0yLjcgMS40LTIuMSAwLTMuNy0xLjUtMy43LTMuN3oiLz48L2c+PHBhdGggZD0ibTE1LjkgMjIuM2M1LjktNC43IDkuOC04LjEgOS44LTExLjQgMC0yLjUtMi0zLjUtMy45LTMuNS0yLjEgMC0zLjguOS00LjcgMi4zbC0xLS45YzEuMi0xLjggMy4zLTIuOCA1LjctMi44IDIuNSAwIDUuNCAxLjQgNS40IDQuOSAwIDMuOC00IDcuMy05IDExLjNoOS4xdjEuM2gtMTEuNHoiLz48L2c+PC9zdmc+">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="2">
			<button type="button" data-matrix-key="ONE" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Im0yMC44IDguMy0yLjggMy0uOS0xIDMuOC00aDEuM3YxNy4zaC0xLjV2LTE1LjN6IiBmaWxsPSIjMDAzODgzIi8+PC9zdmc+">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="3">
			<button type="button" data-matrix-key="EIGHT" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMS44IDMxLjFoLTIuM3YtLjhoNS40di44aC0yLjN2Ni40aC0uOXYtNi40eiIvPjxwYXRoIGQ9Im0xOC4zIDMwLjNoLjl2NC40YzAgMS4zLjcgMi4xIDIgMi4xczItLjggMi0yLjF2LTQuNGguOXY0LjRjMCAxLjgtMSAyLjktMi45IDIuOXMtMi45LTEuMi0yLjktMi45eiIvPjxwYXRoIGQ9Im0yNy4yIDMwLjNoMWwyLjQgNi4yIDIuNC02LjJoMWwtMi45IDcuMmgtMS4xeiIvPjwvZz48cGF0aCBkPSJtMjAuMyAxNC43Yy0yLS41LTQtMS45LTQtNC4yIDAtMy4xIDIuOC00LjUgNS42LTQuNSAyLjcgMCA1LjYgMS40IDUuNiA0LjUgMCAyLjMtMiAzLjYtNCA0LjIgMi4yLjYgNC4zIDIuMiA0LjMgNC42IDAgMi44LTIuNSA0LjYtNS44IDQuNnMtNS45LTEuOC01LjktNC42Yy0uMS0yLjUgMi00LjEgNC4yLTQuNnptMS42LjZjLTEuMS4xLTQuNCAxLjItNC40IDMuOCAwIDIuMSAyLjEgMy40IDQuNCAzLjRzNC40LTEuMyA0LjQtMy40YzAtMi42LTMuNC0zLjYtNC40LTMuOHptMC03LjljLTIuMyAwLTQuMSAxLjItNC4xIDMuMyAwIDIuNCAzLjEgMy4yIDQuMSAzLjQgMS4xLS4yIDQuMS0xIDQuMS0zLjQgMC0yLjEtMS44LTMuMy00LjEtMy4zeiIvPjwvZz48L3N2Zz4=">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="4">
			<button type="button" data-matrix-key="SIX" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMy45IDMxLjYtMi40IDUuOWgtLjRsLTIuNC01Ljl2NS45aC0uOXYtNy4yaDEuM2wyLjIgNS40IDIuMi01LjRoMS4zdjcuMmgtLjl6Ii8+PHBhdGggZD0ibTE5LjUgMzEuOHY1LjdoLS45di03LjJoLjlsNC4xIDUuNnYtNS42aC45djcuMmgtLjl6Ii8+PHBhdGggZD0ibTMxLjcgMzAuMmMyLjEgMCAzLjYgMS42IDMuNiAzLjdzLTEuNCAzLjctMy42IDMuN2MtMi4xIDAtMy42LTEuNi0zLjYtMy43czEuNC0zLjcgMy42LTMuN3ptMCAuOGMtMS43IDAtMi43IDEuMi0yLjcgMi45czEgMi45IDIuNiAyLjkgMi42LTEuMiAyLjYtMi45Yy4xLTEuNy0uOS0yLjktMi41LTIuOXoiLz48L2c+PHBhdGggZD0ibTIyLjYgNmMyLjMgMCAzLjYuOSA0LjcgMi4ybC0uOSAxLjFjLS44LTEuMS0xLjktMS45LTMuOC0xLjktMy43IDAtNS4xIDMuOS01LjEgNy42di44Yy43LTEuMiAyLjctMi45IDUtMi45IDMuMSAwIDUuNiAxLjggNS42IDUuNSAwIDIuOC0yLjEgNS41LTUuOCA1LjUtNC43IDAtNi4zLTQuMy02LjMtOC45IDAtNC41IDEuOC05IDYuNi05em0tLjMgOC4yYy0xLjkgMC0zLjcgMS4yLTQuNyAzIC4yIDIuNCAxLjQgNS40IDQuNyA1LjQgMyAwIDQuMy0yLjMgNC4zLTQuMSAwLTIuOS0xLjgtNC4zLTQuMy00LjN6Ii8+PC9nPjwvc3ZnPg==">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="5">
			<button type="button" data-matrix-key="SEVEN" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im01IDMwLjRoMi45YzEuNCAwIDIuMiAxIDIuMiAyLjJzLS44IDIuMi0yLjIgMi4yaC0ydjIuOWgtLjl6bTIuOC44aC0xLjl2Mi44aDEuOWMuOSAwIDEuNC0uNiAxLjQtMS40cy0uNS0xLjQtMS40LTEuNHoiLz48cGF0aCBkPSJtMTkuMyAzNi43LjcuNy0uNi41LS43LS43Yy0uNS4zLTEuMi41LTEuOS41LTIuMSAwLTMuNi0xLjYtMy42LTMuN3MxLjQtMy43IDMuNi0zLjdjMi4xIDAgMy42IDEuNiAzLjYgMy43LS4xIDEuMS0uNCAyLTEuMSAyLjd6bS0xLjItLjEtMS0xLjEuNi0uNSAxIDEuMWMuNC0uNS43LTEuMi43LTIgMC0xLjctMS0yLjktMi42LTIuOXMtMi42IDEuMi0yLjYgMi45IDEgMi45IDIuNiAyLjljLjUtLjEuOS0uMiAxLjMtLjR6Ii8+PHBhdGggZD0ibTI2LjIgMzQuOGgtMS40djIuOWgtLjl2LTcuMmgyLjljMS4zIDAgMi4yLjggMi4yIDIuMiAwIDEuMy0uOSAyLTEuOSAyLjFsMS45IDIuOWgtMXptLjQtMy42aC0xLjl2Mi44aDEuOWMuOCAwIDEuNC0uNiAxLjQtMS40LjEtLjgtLjUtMS40LTEuNC0xLjR6Ii8+PHBhdGggZD0ibTMyLjcgMzUuOWMuNS41IDEuMiAxIDIuMyAxIDEuMyAwIDEuNy0uNyAxLjctMS4yIDAtLjktLjktMS4xLTEuOC0xLjQtMS4yLS4zLTIuNC0uNi0yLjQtMiAwLTEuMiAxLjEtMiAyLjUtMiAxLjEgMCAxLjkuNCAyLjUgMWwtLjcuN2MtLjUtLjYtMS4zLS45LTIuMS0uOS0uOSAwLTEuNS41LTEuNSAxLjEgMCAuNy44LjkgMS43IDEuMiAxLjIuMyAyLjUuNyAyLjUgMi4yIDAgMS0uNyAyLjEtMi42IDIuMS0xLjIgMC0yLjItLjUtMi44LTEuMXoiLz48L2c+PHBhdGggZD0ibTI0LjkgNy42aC05LjV2LTEuM2gxMS4zdjFsLTcuNCAxNi4yaC0xLjZ6Ii8+PC9nPjwvc3ZnPg==">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="6">
			<button type="button" data-matrix-key="NINE" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im03LjYgMzEuNy0xLjYgNS44aC0xbC0yLTcuMmgxbDEuNiA2IDEuNi02aC44bDEuNiA2IDEuNi02aDFsLTIgNy4yaC0xeiIvPjxwYXRoIGQ9Im0xOCAzNC40LTIuMyAzLjFoLTEuMWwyLjgtMy43LTIuNi0zLjVoMS4xbDIuMSAyLjkgMi4xLTIuOWgxLjFsLTIuNiAzLjUgMi44IDMuN2gtMS4xeiIvPjxwYXRoIGQ9Im0yNi42IDM0LjUtMi44LTQuMWgxbDIuMiAzLjMgMi4yLTMuM2gxbC0yLjggNC4xdjNoLS45di0zeiIvPjxwYXRoIGQ9Im0zMy4xIDM2LjggNC01LjZoLTR2LS44aDUuMnYuN2wtNCA1LjZoNC4xdi44aC01LjJ2LS43eiIvPjwvZz48cGF0aCBkPSJtMTcuNyAyMC42Yy44IDEuMSAxLjkgMS45IDMuOCAxLjkgMy44IDAgNS4xLTQgNS4xLTcuNnYtLjhjLS44IDEuMi0yLjcgMi45LTUuMSAyLjktMy4xIDAtNS42LTEuOC01LjYtNS41LjEtMi44IDIuMi01LjUgNS45LTUuNSA0LjcgMCA2LjMgNC4zIDYuMyA4LjkgMCA0LjQtMS44IDguOS02LjYgOC45LTIuMyAwLTMuNi0uOS00LjYtMi4yem00LjEtMTMuMmMtMyAwLTQuMyAyLjMtNC4zIDQuMSAwIDIuOCAxLjkgNC4yIDQuMyA0LjIgMS45IDAgMy43LTEuMiA0LjctMy0uMi0yLjMtMS40LTUuMy00LjctNS4zeiIvPjwvZz48L3N2Zz4=">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="7">
			<button type="button" data-matrix-key="FOUR" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMy42IDMwLjJjMS4zIDAgMi4yLjYgMi44IDEuM2wtLjcuNWMtLjUtLjYtMS4yLTEtMi4xLTEtMS42IDAtMi44IDEuMi0yLjggMi45czEuMiAyLjkgMi44IDIuOWMuOSAwIDEuNi0uNCAxLjktLjh2LTEuNWgtMi41di0uOGgzLjR2Mi42Yy0uNy43LTEuNiAxLjItMi44IDEuMi0yIDAtMy43LTEuNS0zLjctMy43czEuNy0zLjYgMy43LTMuNnoiLz48cGF0aCBkPSJtMjUuMSAzNC4yaC00LjJ2My4zaC0uOXYtNy4yaC45djMuMWg0LjJ2LTMuMWguOXY3LjJoLS45eiIvPjxwYXRoIGQ9Im0yOS44IDMwLjNoLjl2Ny4yaC0uOXoiLz48L2c+PHBhdGggZD0ibTIzLjYgMTguOGgtOC4ydi0xLjNsNy43LTExLjJoMnYxMS4yaDIuNXYxLjNoLTIuNXY0LjdoLTEuNXptLTYuNy0xLjNoNi43di05Ljd6Ii8+PC9nPjwvc3ZnPg==">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="8">
			<button type="button" data-matrix-key="THREE" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9IiMwMDM4ODMiPjxnIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXciPjxwYXRoIGQ9Im0xMC4yIDMwLjNoMi41YzIuMiAwIDMuNyAxLjYgMy43IDMuNnMtMS41IDMuNi0zLjcgMy42aC0yLjV6bTIuNSA2LjRjMS43IDAgMi44LTEuMiAyLjgtMi44IDAtMS41LTEtMi44LTIuOC0yLjhoLTEuNnY1LjZ6Ii8+PHBhdGggZD0ibTE5LjkgMzAuM2g0Ljd2LjhoLTMuOHYyLjNoMy43di44aC0zLjd2Mi41aDMuOHYuOGgtNC43eiIvPjxwYXRoIGQ9Im0yOC4xIDMwLjNoNC43di44aC0zLjh2Mi4zaDMuN3YuOGgtMy43djMuM2gtLjl6Ii8+PC9nPjxwYXRoIGQ9Im0xNi4zIDIwLjFjMSAxLjQgMi42IDIuNCA0LjggMi40IDIuNyAwIDQuMy0xLjQgNC4zLTMuNyAwLTIuNS0yLTMuNS00LjYtMy41LS43IDAtMS4zIDAtMS42IDB2LTEuM2gxLjZjMi4zIDAgNC40LTEgNC40LTMuMyAwLTIuMS0xLjktMy4zLTQuMS0zLjMtMiAwLTMuNC44LTQuNiAyLjJsLS45LS45YzEuMi0xLjUgMy4xLTIuNyA1LjYtMi43IDMgMCA1LjYgMS42IDUuNiA0LjYgMCAyLjYtMi4yIDMuOC0zLjcgNCAxLjUuMiA0IDEuNCA0IDQuM3MtMi4xIDQuOS01LjggNC45Yy0yLjggMC00LjktMS4zLTUuOS0yLjl6Ii8+PC9nPjwvc3ZnPg==">
			</button>
		</li>
		<li data-matrix-list-item data-matrix-list-item-index="9">
			<button type="button" data-matrix-key="ZERO" class="sasmap__key">
				<img alt="" class="sasmap__img" src="data:image/svg+xml;base64, PHN2ZyBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MiA0MiIgdmlld0JveD0iMCAwIDQyIDQyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Im0yMS41IDZjNC42IDAgNi40IDQuOCA2LjQgOC45cy0xLjggOC45LTYuNCA4LjljLTQuNyAwLTYuNC00LjgtNi40LTguOXMxLjgtOC45IDYuNC04Ljl6bTAgMS40Yy0zLjYgMC00LjggNC00LjggNy42IDAgMy41IDEuMiA3LjYgNC44IDcuNnM0LjgtNCA0LjgtNy42LTEuMi03LjYtNC44LTcuNnoiIGZpbGw9IiMwMDM4ODMiLz48L3N2Zz4=">
			</button>
		</li>
	</ul>
`

const transactionsFirstHTML = `
	<ul class="list list__movement " data-operations-lazy-load="true" data-operations-list="">
		<li class="list-operation-date-line">28 juillet 2022</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-id="id1" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto">
						<span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-no-categorized" style="background-color: #869db8"></span>
					</div>
					<div class="list-operation-item__multiselection">
						<input title="Afficher les détails de l'opération" aria-describedby="movement-id1" class="list-operation-item__multiselection-toggle" id="multiselection-id1" type="checkbox" data-brs-detail-operation-multiple="id1">
						<span class="list-operation-item__fake-checkbox"></span>
					</div>
					<div class="list-operation-item__label" id="movement-id1">
						<div class="list-operation-item__label-name">
							<span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">VIR SEPA MR OU MME CHOU</span>
						</div>
						<div class="list-operation-item__label-sub">
							<span class="list-operation-item__category" title="Non catégorisé">Non catégorisé</span>
						</div>
					</div>
					<div class="list-operation-item__amount positive ">
						1&nbsp;230,29&nbsp;€
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>		
		<li id="movementNativeAdSlot">
			<div id="customer_budget_movements" class="np6-ad " data-brs-np6-slot="" data-brs-np6-slot-xs-id="slot" data-native="1"></div>
		</li>
		<li class="list-operation-item" data-no-select="true" data-id="id2" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__split-picto"><span class="c-icon c-icon--operation-cut"></span></div>
					<div class="list-operation-item__label" id="movement-id2">
						<div class="list-operation-item__label-name"><span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">CARTE 21/11/21 MR POU CB*2100</span></div>
						<div class="list-operation-item__label-sub"></div>
					</div>
					<div class="list-operation-item__amount positive ">
                    	50,99&nbsp;€
                	</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
			<ul class="list list__movement ">
				<li class="list-operation-item list-operation-item--splitted" data-brs-detail-operation="" data-id="id2-1" data-parent-id="id2" data-child="true" data-operations-item="">
					<div class="list-operation-item__block list__movement__line--block__split">
						<div class="list-operation-item__content">
							<div class="list-operation-item__category-picto"><span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-insurance" style="background-color: #E5C010"></span></div>
							<div class="list-operation-item__multiselection">
								<input title="Afficher les détails de l'opération" aria-describedby="movement-id2-1" class="list-operation-item__multiselection-toggle" id="multiselection-id2-1" type="checkbox" data-brs-detail-operation-multiple="id2-1">
								<span class="list-operation-item__fake-checkbox"></span>
							</div>
							<div class="list-operation-item__label" id="movement-id2-1">
								<div class="list-operation-item__label-name"><span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">CARTE 21/11/21 MR POU CB*2100</span></div>
								<div class="list-operation-item__label-sub"><span class="list-operation-item__category" title="Carburant">Carburant</span></div>
							</div>
							<div class="list-operation-item__amount positive ">
								49,99&nbsp;€
							</div>
							<div class="list-operation-item__check">
								<div class="list-operation-item__check-button u-force-focus " data-account-key="id2-1" data-movement-id="id2-1" data-action="check" data-check-operation="" role="checkbox" tabindex="0" aria-describedby="movement-id2-1" aria-checked="false" aria-label="Pointer l'opération" title="Pointer l'opération"><span class="c-icon c-icon--check"></span></div>
							</div>
						</div>
					</div>
				</li>
				<li class="list-operation-item list-operation-item--splitted" data-brs-detail-operation="" data-id="id2-2" data-parent-id="id2" data-child="true" data-operations-item="">
					<div class="list-operation-item__block list__movement__line--block__split">
						<div class="list-operation-item__content">
							<div class="list-operation-item__category-picto"><span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-insurance" style="background-color: #E5C010"></span></div>
							<div class="list-operation-item__multiselection">
								<input title="Afficher les détails de l'opération" aria-describedby="movement-id2-2" class="list-operation-item__multiselection-toggle" id="multiselection-id2-2" type="checkbox" data-brs-detail-operation-multiple="id2-2">
								<span class="list-operation-item__fake-checkbox"></span>
							</div>
							<div class="list-operation-item__label" id="movement-id2-2">
								<div class="list-operation-item__label-name"><span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">Splitted</span></div>
								<div class="list-operation-item__label-sub"><span class="list-operation-item__category" title="Auto">Auto</span></div>
							</div>
							<div class="list-operation-item__amount positive ">
								10,00&nbsp;€
							</div>
							<div class="list-operation-item__check">
								<div class="list-operation-item__check-button u-force-focus " data-account-key="id2-2" data-movement-id="id2-2" data-action="check" data-check-operation="" role="checkbox" tabindex="0" aria-describedby="movement-id2-2" aria-checked="false" aria-label="Pointer l'opération" title="Pointer l'opération"><span class="c-icon c-icon--check"></span></div>
							</div>
						</div>
					</div>
				</li>
			</ul>
		</li>
		<li class="list-operation-date-line">22 janvier 2022</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-id="id3" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto">
						<span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-no-categorized" style="background-color: #869db8"></span>
					</div>
					<div class="list-operation-item__multiselection">
						<input title="Afficher les détails de l'opération" aria-describedby="movement-id3" class="list-operation-item__multiselection-toggle" id="multiselection-id3" type="checkbox" data-brs-detail-operation-multiple="id3">
						<span class="list-operation-item__fake-checkbox"></span>
					</div>
					<div class="list-operation-item__label" id="movement-id3">
						<div class="list-operation-item__label-name">
							<span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">PRLV SEPA MR OU MME CHOUPETTE</span>
						</div>
						<div class="list-operation-item__label-sub">
							<span class="list-operation-item__category" title="Catégorie 2">Catégorie 2</span>
						</div>
					</div>
					<div class="list-operation-item__amount neutral ">
						− 20,49 € 
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>
		<li class="list__movement__range-summary list__movement__range-summary--loading-placeholder list__movement__range-summary--loading-trigger" data-operations-next-pagination="next-page-1" data-operations-date-range="">
			<div data-operations-loading-title="" tabindex="-1" class="list__movement__range-summary__loading-title u-text-center">
        		Récupération des mouvements ...
        		<div class="bouncy-loader ">
					<div class="bouncy-loader__balls">
						<div class="bouncy-loader__ball bouncy-loader__ball--left"></div>
						<div class="bouncy-loader__ball bouncy-loader__ball--center"></div>
						<div class="bouncy-loader__ball bouncy-loader__ball--right"></div>
					</div>
				</div>
			</div>
			<a data-operations-next-pagination-trigger="" href="javascript://;" class="list__movement__range-summary__loading-trigger u-text-center" data-tag-commander-click="{&quot;label&quot; : &quot;mes-comptes::mes-mouvements::mouvements-suivants&quot;, &quot;s2&quot;: 1, &quot;type&quot; : &quot;N&quot;}">
        		Mouvements précédents
    		</a>
		</li>	
	</ul>
`

const transactionsSecondHTML = `
	<ul class="list list__movement " data-operations-lazy-load="true" data-operations-list="">
		<li class="list-operation-date-line">21 août 2021</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-external="true" data-id="id4" data-tag-commander-click="pending_authorizations" data-account-key="id4" data-is-auth="true" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto--authorization"><span class="c-icon--authorization2" style="color: #3d5d84"></span></div>
					<div class="list-operation-item__label" id="movement-id4">
						<div class="list-operation-item__label-name--authorization"><span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">POULET</span></div>
						<div class="list-operation-item__label-sub"><span class="list-operation-item__category" title="Autorisation paiement en cours CB">Autorisation paiement en cours CB*2100</span></div>
					</div>
					<div class="list-operation-item__amount list-operation-item__amount--authorization ">
						−&nbsp;19,16&nbsp;€
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>
		<li class="list-operation-date-line">22 janvier 2021</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-id="id5" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto">
						<span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-no-categorized" style="background-color: #869db8"></span>
					</div>
					<div class="list-operation-item__multiselection">
						<input title="Afficher les détails de l'opération" aria-describedby="movement-id5" class="list-operation-item__multiselection-toggle" id="multiselection-id5" type="checkbox" data-brs-detail-operation-multiple="id5">
						<span class="list-operation-item__fake-checkbox"></span>
					</div>
					<div class="list-operation-item__label" id="movement-id5">
						<div class="list-operation-item__label-name">
							<span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">VIR INST MR OU MME CHOUPETTEL</span>
						</div>
						<div class="list-operation-item__label-sub">
							<span class="list-operation-item__category" title="Catégorie 5">Catégorie 5</span>
						</div>
					</div>
					<div class="list-operation-item__amount neutral ">
						− 10,11 € 
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-id="id9" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto">
						<span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-no-categorized" style="background-color: #869db8"></span>
					</div>
					<div class="list-operation-item__multiselection">
						<input title="Afficher les détails de l'opération" aria-describedby="movement-id9" class="list-operation-item__multiselection-toggle" id="multiselection-id9" type="checkbox" data-brs-detail-operation-multiple="id9">
						<span class="list-operation-item__fake-checkbox"></span>
					</div>
					<div class="list-operation-item__label" id="movement-id9">
						<div class="list-operation-item__label-name">
							<span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">VIR Corner - Amazoune</span>
						</div>
						<div class="list-operation-item__label-sub">
							<span class="list-operation-item__category" title="Catégorie 1">Catégorie 1</span>
						</div>
					</div>
					<div class="list-operation-item__amount neutral ">
						− 101,41 € 
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>
		<li class="list-operation-item" data-brs-detail-operation="" data-id="id2" data-is-auth="false" data-operations-item="">
			<div class="list-operation-item__block ">
				<div class="list-operation-item__content">
					<div class="list-operation-item__category-picto">
						<span class="c-icon c-icon--s c-icon--rounded c-icon--color-white c-icon--pfm-no-categorized" style="background-color: #869db8"></span>
					</div>
					<div class="list-operation-item__multiselection">
						<input title="Afficher les détails de l'opération" aria-describedby="movement-id2" class="list-operation-item__multiselection-toggle" id="multiselection-id2" type="checkbox" data-brs-detail-operation-multiple="id2">
						<span class="list-operation-item__fake-checkbox"></span>
					</div>
					<div class="list-operation-item__label" id="movement-id2">
						<div class="list-operation-item__label-name">
							<span title="Afficher les détails de l'opération" class="list__movement--label-user" data-operation-label-long="">AVOIR 01/02/19 BOUBI BOUBI CB*29</span>
						</div>
						<div class="list-operation-item__label-sub">
							<span class="list-operation-item__category" title="Catégorie 2">Catégorie 2</span>
						</div>
					</div>
					<div class="list-operation-item__amount neutral ">
						− 9,21 € 
					</div>
					<div class="list-operation-item__check"></div>
				</div>
			</div>
		</li>
		<li class="list__movement__range-summary list__movement__range-summary--loading-placeholder list__movement__range-summary--loading-trigger">
			<div data-operations-loading-title="" tabindex="-1" class="list__movement__range-summary__loading-title u-text-center">
				Récupération des mouvements ...
				<div class="bouncy-loader ">
					<div class="bouncy-loader__balls">
						<div class="bouncy-loader__ball bouncy-loader__ball--left"></div>
						<div class="bouncy-loader__ball bouncy-loader__ball--center"></div>
						<div class="bouncy-loader__ball bouncy-loader__ball--right"></div>
					</div>
				</div>
			</div>
		</li>
	</ul>
`
