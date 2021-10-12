let step = 0;
let pivots = [1, 3, 3, 1];

function rand(min, max) {
  // min and max included
  return Math.floor(Math.random() * (max - min + 1) + min);
}

function quickSort(arr, left, right) {
  let index;

  //We want to check if we even need to do the sorting
  if (arr.length > 1) {
    //We first sort the function, and return the center which is used to divide the array
    index = sortingUsingPivot(arr, left, right);

    //If there are more elements on the left side of pivot that needs to be sorted
    if (left < index - 1) {
      //will put the entire left of the array into the quicksort again
      quickSort(arr, left, index - 1);
    }

    //If there are more elements on the right side of pivot that needs to be sorted
    if (index < right) {
      //will put the entire right of the array into the quicksort again
      quickSort(arr, index, right);
    }
  }
}

function sortingUsingPivot(arr, left, right) {
  pivotIndex = pivots[step];
  // pivotIndex = left + rand(0, right - left);
  //We are using the middle element of the array as our pivot
  /* let pivot = arr[Math.floor((right + left) / 2)] */
  let pivot = arr[pivotIndex];
  let l = left; //This is keeping track of left pointer
  let r = right; //this is keeping track of right pointer
  console.log("pivot", pivot, "pivotIndex", pivotIndex);
  console.log("i", l, "j", r);
  console.log("START");
  console.log(arr.slice(left, right + 1));

  //Keep going until left pointer passes the right pointer
  while (l <= r) {
    //Used Find the first element on the left side that is larger than the pivot element.
    //So that we know this is the element we want to move to the other side
    while (arr[l] < pivot) {
      //Keep searching until we pass the pivot
      l++;
      console.log("i++", l);
    }

    //Used Find the first element on the right side that is smaller than the pivot element.
    //So that we know this is the element we want to move to the other side
    while (arr[r] > pivot) {
      //Keep searching until we pass the pivot
      r--;
      console.log("j--", r);
    }

    //we want to swap the two elements as long as left pointer doesn't pass the right pointer
    if (l <= r) {
      console.log("swap", arr[l], "and", arr[r]);
      swap(arr, l, r);
      l++;
      r--;
      console.log("i++", l, "j--", r);
    }
  }

  console.log("FINISH");
  console.log(arr.slice(left, right + 1));
  console.log("");
  step++;
  //Return the left pointer as that is our new center to divide the array
  return l;
}

function swap(arr, leftIndex, rightIndex) {
  //We basically just swap the two items from the two different index
  let temp = arr[leftIndex];
  arr[leftIndex] = arr[rightIndex];
  arr[rightIndex] = temp;
}

arr = [-5, -10, 2, 3, 1, 3];
quickSort(arr, 0, arr.length - 1);
console.log(arr);
