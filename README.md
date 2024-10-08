# [_Memos Notion link_](https://quaint-fascinator-7a3.notion.site/Xu-leetcode-Summary-ef3bd2078fda4d8281a406e8989215af?pvs=4)

---

## Setup Script:

**Source Code**:
[leetcode-setup.go](leetcode-setup.go)

**Download**:

- windows: [leetcode-setup.exe](leetcode-setup.exe)

```bash
wget https://github.com/comevback/leetcode/blob/main/leetcode-setup.exe -O leetcode-setup.exe
```

- macOS: [leetcode-setup-mac](leetcode-setup-mac)

```bash
wget https://github.com/comevback/leetcode/blob/main/leetcode-setup-mac -O leetcode-setup
```

- linux: [leetcode-setup-linux](leetcode-setup-linux)

```bash
wget https://github.com/comevback/leetcode/blob/main/leetcode-setup-linux -O leetcode-setup
```

### Setup as a Global Command:

#### For macOS and Linux Users

Move the Binary File and Set Execution Permissions:

```bash
sudo mv leetcode-setup /usr/local/bin/ &&
sudo chmod +x /usr/local/bin/leetcode-setup
```

---

#### For Windows Users

1. **Create a Directory:**

   - Create a directory where you will store the binary. Recommended locations are `C:\Tools` or `C:\Users\YourUsername\bin`.

2. **Move the Binary File:**

   - Move the `leetcode-setup.exe` file into the directory you created.

3. **Add the Directory to the PATH:**
   - Right-click on **"This PC"** or **"My Computer"**, and select **"Properties"**.
   - Click on **"Advanced system settings"**.
   - In the **"System Properties"** window, click on **"Environment Variables"**.
   - In the **"System variables"** section, find and select the **"Path"** variable, then click on **"Edit"**.
   - In the **"Edit environment variable"** window, click **"New"** and enter the path to your directory, e.g., `C:\Tools`.
   - Click **"OK"** to save the changes.

---

## Fundamentals:

[Go-Packages.md](Go-Packages.md)

[BigO-cheat-sheet.md](BigO-cheat-sheet.md)

[regex-cheatsheet.pdf](regex-cheatsheet.pdf)

[CustomizedSplit.md](Paiza/space-input/CustomizedSplit.md)

[System-Design-Template.md](System-Design-Template.md)

[LinkedList-implementation.md](DSA/LinkedList/LinkedList-implementation.md)

[Stack-implementation.md](DSA/Stack-Queue/Stack/Stack-implementation.md)

[Queue-implementation.md](DSA/Stack-Queue/Queue/Queue-implementation.md)

[MaxMinQueue-implementation.md](DSA/Stack-Queue/MaxMinQueue/MaxMinQueue.md)

[MaxMinStack-implementation.md](DSA/Stack-Queue/MaxMinStack/MaxMinStack.md)

[Priority-Queue-implementation.md](DSA/Stack-Queue/Priority-Queue/Priority-Queue.md)

[Binary-Tree.md](DSA/Binary-Tree/Binary-Tree.md)

[KMP.md](DSA/KMP/KMP.md)

[Recursion.md](DSA/Recursion/Recursion.md)

---

## Sorting:

[692-Top-K-Frequent-Words.md](Code/692-Top-K-Frequent-Words/692-Top-K-Frequent-Words.md)

[Selection-Sort.md](DSA/Sorting/Selection-Sort/Selection-Sort.md)

[Bubble-Sort.md](DSA/Sorting/Bubble-Sort/Bubble-Sort.md)

[Insertion-Sort.md](DSA/Sorting/Insertion-Sort/Insertion-Sort.md)

[Merge-Sort.md](DSA/Sorting/Merge-Sort/Merge-Sort.md)

[Quick-Sort.md](DSA/Sorting/Quick-Sort/Quick-Sort.md)

[Binary-Search.md](DSA/Binary-Search/Binary-Search.md)

---

## Binary-Tree 100

1. [104-Maximum-Depth-of-Binary-Tree.md](Binary-Tree/104-Maximum-Depth-of-Binary-Tree/104-Maximum-Depth-of-Binary-Tree.md)
2. [LCR-175-binary-tree-depth.md](Binary-Tree/LCR-175-binary-tree-depth/LCR-175-binary-tree-depth.md)
3. [144-Binary-Tree-Preorder-Traversal.md](Binary-Tree/144-Binary-Tree-Preorder-Traversal/144-Binary-Tree-Preorder-Traversal.md)
4. [543-Diameter-of-Binary-Tree.md](Binary-Tree/543-Diameter-of-Binary-Tree/543-Diameter-of-Binary-Tree.md)
5. [257-Binary-Tree-Paths.md](Binary-Tree/257-Binary-Tree-Paths/257-Binary-Tree-Paths.md)
6. [226-Invert-Binary-Tree.md](Binary-Tree/226-Invert-Binary-Tree/226-Invert-Binary-Tree.md)
7. [116-Populating-Next-Right-Pointers-in-Each-Node.md](Binary-Tree/116-Populating-Next-Right-Pointers-in-Each-Node/116-Populating-Next-Right-Pointers-in-Each-Node.md)
8. [114-Flatten-Binary-Tree-to-Linked-List.md](Binary-Tree/114-Flatten-Binary-Tree-to-Linked-List/114-Flatten-Binary-Tree-to-Linked-List.md)
9. [654-Maximum-Binary-Tree.md](Binary-Tree/654-Maximum-Binary-Tree/654-Maximum-Binary-Tree.md)
10. [105-Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal.md](Binary-Tree/105-Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal/105-Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal.md)
11. [106-Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal.md](Binary-Tree/106-Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal/106-Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal.md)
12. [889-Construct-Binary-Tree-from-Preorder-and-Postorder-Traversal.md](Binary-Tree/889-Construct-Binary-Tree-from-Preorder-and-Postorder-Traversal/889-Construct-Binary-Tree-from-Preorder-and-Postorder-Traversal.md)
13. [652-Find-Duplicate-Subtrees.md](Binary-Tree/652-Find-Duplicate-Subtrees/652-Find-Duplicate-Subtrees.md)
14. [297-Serialize-and-Deserialize-Binary-Tree.md](Binary-Tree/297-Serialize-and-Deserialize-Binary-Tree/297-Serialize-and-Deserialize-Binary-Tree.md)
15. [912-Sort-an-Array.md](Binary-Tree/912-Sort-an-Array/912-Sort-an-Array.md)
16. [315-Count-of-Smaller-Numbers-After-Self.md](Binary-Tree/nf-315-Count-of-Smaller-Numbers-After-Self/315-Count-of-Smaller-Numbers-After-Self.md)
17. [1038-Binary-Search-Tree-to-Greater-Sum-Tree.md](Binary-Tree/1038-Binary-Search-Tree-to-Greater-Sum-Tree/1038-Binary-Search-Tree-to-Greater-Sum-Tree.md)
18. [538-Convert-BST-to-Greater-Tree.md](Binary-Tree/538-Convert-BST-to-Greater-Tree/538-Convert-BST-to-Greater-Tree.md)
19. [701-Insert-into-a-Binary-Search-Tree.md](Binary-Tree/701-Insert-into-a-Binary-Search-Tree/701-Insert-into-a-Binary-Search-Tree.md)
20. [98-Validate-Binary-Search-Tree.md](Binary-Tree/98-Validate-Binary-Search-Tree/98-Validate-Binary-Search-Tree.md)
21. [95-Unique-Binary-Search-Trees-II.md](Binary-Tree/95-Unique-Binary-Search-Trees-II/95-Unique-Binary-Search-Trees-II.md)
22. [96-Unique-Binary-Search-Trees.md](Binary-Tree/96-Unique-Binary-Search-Trees/96-Unique-Binary-Search-Trees.md)
23. [215-Kth-Largest-Element-in-an-Array.md](Binary-Tree/215-Kth-Largest-Element-in-an-Array/215-Kth-Largest-Element-in-an-Array.md)
24. [236-Lowest-Common-Ancestor-of-a-Binary-Tree.md](Binary-Tree/236-Lowest-Common-Ancestor-of-a-Binary-Tree/236-Lowest-Common-Ancestor-of-a-Binary-Tree.md)
25. [235-Lowest-Common-Ancestor-of-a-Binary-Search-Tree.md](Binary-Tree/235-Lowest-Common-Ancestor-of-a-Binary-Search-Tree/235-Lowest-Common-Ancestor-of-a-Binary-Search-Tree.md)
26. [222-Count-Complete-Tree-Nodes.md](Binary-Tree/222-Count-Complete-Tree-Nodes/222-Count-Complete-Tree-Nodes.md)
27. [199-Binary-Tree-Right-Side-View.md](Binary-Tree/199-Binary-Tree-Right-Side-View/199-Binary-Tree-Right-Side-View.md)
28. [298-binary-tree-longest-consecutive-sequence.md](Binary-Tree/298-binary-tree-longest-consecutive-sequence/298-binary-tree-longest-consecutive-sequence.md)
29. [988-Smallest-String-Starting-From-Leaf.md](Binary-Tree/988-Smallest-String-Starting-From-Leaf/988-Smallest-String-Starting-From-Leaf.md)
30. [1457-Pseudo-Palindromic-Paths-in-a-Binary-Tree.md](Binary-Tree/1457-Pseudo-Palindromic-Paths-in-a-Binary-Tree/1457-Pseudo-Palindromic-Paths-in-a-Binary-Tree.md)
31. [270-closest-binary-search-tree-value.md](Binary-Tree/270-closest-binary-search-tree-value/270-closest-binary-search-tree-value.md)
32. [404-Sum-of-Left-Leaves.md](Binary-Tree/404-Sum-of-Left-Leaves/404-Sum-of-Left-Leaves.md)
33. [623-Add-One-Row-to-Tree.md](Binary-Tree/623-Add-One-Row-to-Tree/623-Add-One-Row-to-Tree.md)
34. [971-Flip-Binary-Tree-To-Match-Preorder-Traversal.md](Binary-Tree/971-Flip-Binary-Tree-To-Match-Preorder-Traversal/971-Flip-Binary-Tree-To-Match-Preorder-Traversal.md)
35. [987-Vertical-Order-Traversal-of-a-Binary-Tree.md](Binary-Tree/987-Vertical-Order-Traversal-of-a-Binary-Tree/987-Vertical-Order-Traversal-of-a-Binary-Tree.md)
36. [993-Cousins-in-Binary-Tree.md](Binary-Tree/993-Cousins-in-Binary-Tree/993-Cousins-in-Binary-Tree.md)
37. [1315-Sum-of-Nodes-with-Even-Valued-Grandparent.md](Binary-Tree/1315-Sum-of-Nodes-with-Even-Valued-Grandparent/1315-Sum-of-Nodes-with-Even-Valued-Grandparent.md)
38. [1448-Count-Good-Nodes-in-Binary-Tree.md](Binary-Tree/1448-Count-Good-Nodes-in-Binary-Tree/1448-Count-Good-Nodes-in-Binary-Tree.md)
39. [1469-find-all-the-lonely-nodes.md](Binary-Tree/1469-find-all-the-lonely-nodes/1469-find-all-the-lonely-nodes.md)
40. [1602-find-nearest-right-node-in-binary-tree.md](Binary-Tree/1602-find-nearest-right-node-in-binary-tree/1602-find-nearest-right-node-in-binary-tree.md)
41. [437-Path-Sum-III.md](Binary-Tree/437-Path-Sum-III/437-Path-Sum-III.md)
42. [513-Find-Bottom-Left-Tree-Value.md](Binary-Tree/513-Find-Bottom-Left-Tree-Value/513-Find-Bottom-Left-Tree-Value.md)
43. [1261-Find-Elements-in-a-Contaminated-Binary-Tree.md](Binary-Tree/1261-Find-Elements-in-a-Contaminated-Binary-Tree/1261-Find-Elements-in-a-Contaminated-Binary-Tree.md)
44. [100-Same-Tree.md](Binary-Tree/100-Same-Tree/100-Same-Tree.md)
45. [572-Subtree-of-Another-Tree.md](Binary-Tree/572-Subtree-of-Another-Tree/572-Subtree-of-Another-Tree.md)
46. [1367-Linked-List-in-Binary-Tree.md](Binary-Tree/1367-Linked-List-in-Binary-Tree/1367-Linked-List-in-Binary-Tree.md)
47. [894-All-Possible-Full-Binary-Trees.md](Binary-Tree/894-All-Possible-Full-Binary-Trees/894-All-Possible-Full-Binary-Trees.md)
48. [998-Maximum-Binary-Tree-II.md](Binary-Tree/998-Maximum-Binary-Tree-II/998-Maximum-Binary-Tree-II.md)
49. [1110-Delete-Nodes-And-Return-Forest.md](Binary-Tree/1110-Delete-Nodes-And-Return-Forest/1110-Delete-Nodes-And-Return-Forest.md)
50. [101-Symmetric-Tree.md](Binary-Tree/101-Symmetric-Tree/101-Symmetric-Tree.md)
51. [951-Flip-Equivalent-Binary-Trees.md](Binary-Tree/951-Flip-Equivalent-Binary-Trees/951-Flip-Equivalent-Binary-Trees.md)
52. [124-Binary-Tree-Maximum-Path-Sum.md](Binary-Tree/124-Binary-Tree-Maximum-Path-Sum/124-Binary-Tree-Maximum-Path-Sum.md)
53. [112-Path-Sum.md](Binary-Tree/112-Path-Sum/112-Path-Sum.md)
54. [113-Path-Sum-II.md](Binary-Tree/113-Path-Sum-II/113-Path-Sum-II.md)
55. [617-Merge-Two-Binary-Trees.md](Binary-Tree/617-Merge-Two-Binary-Trees/617-Merge-Two-Binary-Trees.md)
56. [897-Increasing-Order-Search-Tree.md](Binary-Tree/897-Increasing-Order-Search-Tree/897-Increasing-Order-Search-Tree.md)
57. [938-Range-Sum-of-BST.md](Binary-Tree/938-Range-Sum-of-BST/938-Range-Sum-of-BST.md)
58. [1379-Find-a-Corresponding-Node-of-a-Binary-Tree-in-a-Clone-of-That-Tree.md](Binary-Tree/1379-Find-a-Corresponding-Node-of-a-Binary-Tree-in-a-Clone-of-That-Tree/1379-Find-a-Corresponding-Node-of-a-Binary-Tree-in-a-Clone-of-That-Tree.md)
59. [110-Balanced-Binary-Tree.md](Binary-Tree/110-Balanced-Binary-Tree/110-Balanced-Binary-Tree.md)
60. [508-Most-Frequent-Subtree-Sum.md](Binary-Tree/508-Most-Frequent-Subtree-Sum/508-Most-Frequent-Subtree-Sum.md)
61. [563-Binary-Tree-Tilt.md](Binary-Tree/563-Binary-Tree-Tilt/563-Binary-Tree-Tilt.md)
62. [687-Longest-Univalue-Path.md](Binary-Tree/687-Longest-Univalue-Path/687-Longest-Univalue-Path.md)
63. [865-Smallest-Subtree-with-all-the-Deepest-Nodes.md](Binary-Tree/865-Smallest-Subtree-with-all-the-Deepest-Nodes/865-Smallest-Subtree-with-all-the-Deepest-Nodes.md)
64. [1026-Maximum-Difference-Between-Node-and-Ancestor.md](Binary-Tree/1026-Maximum-Difference-Between-Node-and-Ancestor/1026-Maximum-Difference-Between-Node-and-Ancestor.md)
65. [1339-Maximum-Product-of-Splitted-Binary-Tree.md](Binary-Tree/1339-Maximum-Product-of-Splitted-Binary-Tree/1339-Maximum-Product-of-Splitted-Binary-Tree.md)
66. [1372-Longest-ZigZag-Path-in-a-Binary-Tree.md](Binary-Tree/1372-Longest-ZigZag-Path-in-a-Binary-Tree/1372-Longest-ZigZag-Path-in-a-Binary-Tree.md)
67. [814-Binary-Tree-Pruning.md](Binary-Tree/814-Binary-Tree-Pruning/814-Binary-Tree-Pruning.md)
68. [2049-Count-Nodes-With-the-Highest-Score.md](Binary-Tree/2049-Count-Nodes-With-the-Highest-Score/2049-Count-Nodes-With-the-Highest-Score.md)
69. [968-Binary-Tree-Cameras.md](Binary-Tree/968-Binary-Tree-Cameras/968-Binary-Tree-Cameras.md)
70. [979-Distribute-Coins-in-Binary-Tree.md](Binary-Tree/979-Distribute-Coins-in-Binary-Tree/979-Distribute-Coins-in-Binary-Tree.md)
71. [1080-Insufficient-Nodes-in-Root-to-Leaf-Paths.md](Binary-Tree/1080-Insufficient-Nodes-in-Root-to-Leaf-Paths/1080-Insufficient-Nodes-in-Root-to-Leaf-Paths.md)
72. [117-Populating-Next-Right-Pointers-in-Each-Node-II.md](Binary-Tree/117-Populating-Next-Right-Pointers-in-Each-Node-II/117-Populating-Next-Right-Pointers-in-Each-Node-II.md)
73. [662-Maximum-Width-of-Binary-Tree.md](Binary-Tree/662-Maximum-Width-of-Binary-Tree/662-Maximum-Width-of-Binary-Tree.md)
74. [515-Find-Largest-Value-in-Each-Tree-Row.md](Binary-Tree/515-Find-Largest-Value-in-Each-Tree-Row/515-Find-Largest-Value-in-Each-Tree-Row.md)
75. [637-Average-of-Levels-in-Binary-Tree.md](Binary-Tree/637-Average-of-Levels-in-Binary-Tree/637-Average-of-Levels-in-Binary-Tree.md)
76. [102-Binary-Tree-Level-Order-Traversal.md](Binary-Tree/102-Binary-Tree-Level-Order-Traversal/102-Binary-Tree-Level-Order-Traversal.md)
77. [958-Check-Completeness-of-a-Binary-Tree.md](Binary-Tree/958-Check-Completeness-of-a-Binary-Tree/958-Check-Completeness-of-a-Binary-Tree.md)
78. [1161-Maximum-Level-Sum-of-a-Binary-Tree.md](Binary-Tree/1161-Maximum-Level-Sum-of-a-Binary-Tree/1161-Maximum-Level-Sum-of-a-Binary-Tree.md)
79. [1302-Deepest-Leaves-Sum.md](Binary-Tree/1302-Deepest-Leaves-Sum/1302-Deepest-Leaves-Sum.md)
80. [1609-Even-Odd-Tree.md](Binary-Tree/1609-Even-Odd-Tree/1609-Even-Odd-Tree.md)
81. [103-Binary-Tree-Zigzag-Level-Order-Traversal.md](Binary-Tree/103-Binary-Tree-Zigzag-Level-Order-Traversal/103-Binary-Tree-Zigzag-Level-Order-Traversal.md)
82. [919-Complete-Binary-Tree-Inserter.md](Binary-Tree/919-Complete-Binary-Tree-Inserter/919-Complete-Binary-Tree-Inserter.md)
83. [863-All-Nodes-Distance-K-in-Binary-Tree.md](Binary-Tree/863-All-Nodes-Distance-K-in-Binary-Tree/863-All-Nodes-Distance-K-in-Binary-Tree.md)
84. [582-kill-process.md](Binary-Tree/582-kill-process/582-kill-process.md)
85. [536-construct-binary-tree-from-string.md](Binary-Tree/536-construct-binary-tree-from-string/536-construct-binary-tree-from-string.md)
86. [99-Recover-Binary-Search-Tree.md](Binary-Tree/99-Recover-Binary-Search-Tree/99-Recover-Binary-Search-Tree.md)
87. [669-Trim-a-Binary-Search-Tree.md](Binary-Tree/669-Trim-a-Binary-Search-Tree/669-Trim-a-Binary-Search-Tree.md)
88. [671-Second-Minimum-Node-In-a-Binary-Tree.md](Binary-Tree/671-Second-Minimum-Node-In-a-Binary-Tree/671-Second-Minimum-Node-In-a-Binary-Tree.md)
89. [1008-Construct-Binary-Search-Tree-from-Preorder-Traversal.md](Binary-Tree/1008-Construct-Binary-Search-Tree-from-Preorder-Traversal/1008-Construct-Binary-Search-Tree-from-Preorder-Traversal.md)
90. [108-Convert-Sorted-Array-to-Binary-Search-Tree.md](Binary-Tree/108-Convert-Sorted-Array-to-Binary-Search-Tree/108-Convert-Sorted-Array-to-Binary-Search-Tree.md)
91. [173-Binary-Search-Tree-Iterator.md](Binary-Tree/173-Binary-Search-Tree-Iterator/173-Binary-Search-Tree-Iterator.md)
92. [449-Serialize-and-Deserialize-BST.md](Binary-Tree/449-Serialize-and-Deserialize-BST/449-Serialize-and-Deserialize-BST.md)
93. [1305-All-Elements-in-Two-Binary-Search-Trees.md](Binary-Tree/1305-All-Elements-in-Two-Binary-Search-Trees/1305-All-Elements-in-Two-Binary-Search-Trees.md)
94. [109-Convert-Sorted-List-to-Binary-Search-Tree.md](Binary-Tree/109-Convert-Sorted-List-to-Binary-Search-Tree/109-Convert-Sorted-List-to-Binary-Search-Tree.md)
95. [530-Minimum-Absolute-Difference-in-BST.md](Binary-Tree/530-Minimum-Absolute-Difference-in-BST/530-Minimum-Absolute-Difference-in-BST.md)
96. [129-Sum-Root-to-Leaf-Numbers.md](Binary-Tree/129-Sum-Root-to-Leaf-Numbers/129-Sum-Root-to-Leaf-Numbers.md)
97. [872-Leaf-Similar-Trees.md](Binary-Tree/872-Leaf-Similar-Trees/872-Leaf-Similar-Trees.md)
98. [94-Binary-Tree-Inorder-Traversal.md](Binary-Tree/94-Binary-Tree-Inorder-Traversal/94-Binary-Tree-Inorder-Traversal.md)
99. [111-Minimum-Depth-of-Binary-Tree.md](Binary-Tree/111-Minimum-Depth-of-Binary-Tree/111-Minimum-Depth-of-Binary-Tree.md)
100. [107-Binary-Tree-Level-Order-Traversal-II.md](Binary-Tree/107-Binary-Tree-Level-Order-Traversal-II/107-Binary-Tree-Level-Order-Traversal-II.md)
101. [145-Binary-Tree-Postorder-Traversal.md](Binary-Tree/145-Binary-Tree-Postorder-Traversal/145-Binary-Tree-Postorder-Traversal.md)
102. [501-Find-Mode-in-Binary-Search-Tree.md](Binary-Tree/501-Find-Mode-in-Binary-Search-Tree/501-Find-Mode-in-Binary-Search-Tree.md)
103. [655-Print-Binary-Tree.md](Binary-Tree/655-Print-Binary-Tree/655-Print-Binary-Tree.md)
104. [703-Kth-Largest-Element-in-a-Stream.md](Binary-Tree/703-Kth-Largest-Element-in-a-Stream/703-Kth-Largest-Element-in-a-Stream.md)
105. [2236-Root-Equals-Sum-of-Children.md](Binary-Tree/2236-Root-Equals-Sum-of-Children/2236-Root-Equals-Sum-of-Children.md)
106. [2265-Count-Nodes-Equal-to-Average-of-Subtree.md](Binary-Tree/2265-Count-Nodes-Equal-to-Average-of-Subtree/2265-Count-Nodes-Equal-to-Average-of-Subtree.md)
107. [2331-Evaluate-Boolean-Binary-Tree.md](Binary-Tree/2331-Evaluate-Boolean-Binary-Tree/2331-Evaluate-Boolean-Binary-Tree.md)
108. [2583-Kth-Largest-Sum-in-a-Binary-Tree.md](Binary-Tree/2583-Kth-Largest-Sum-in-a-Binary-Tree/2583-Kth-Largest-Sum-in-a-Binary-Tree.md)
109. [2385-Amount-of-Time-for-Binary-Tree-to-Be-Infected.md](Binary-Tree/2385-Amount-of-Time-for-Binary-Tree-to-Be-Infected/2385-Amount-of-Time-for-Binary-Tree-to-Be-Infected.md)
110. [2415-Reverse-Odd-Levels-of-Binary-Tree.md](Binary-Tree/2415-Reverse-Odd-Levels-of-Binary-Tree/2415-Reverse-Odd-Levels-of-Binary-Tree.md)

---

## Two-Pointer

[86-Partition-List.md](Two-Pointer/86-Partition-List/86-Partition-List.md)

[21-Merge-Two-Sorted-Lists.md](Two-Pointer/21-Merge-Two-Sorted-Lists/21-Merge-Two-Sorted-Lists.md)

[23-Merge-k-Sorted-Lists.md](Two-Pointer/23-Merge-k-Sorted-Lists/23-Merge-k-Sorted-Lists.md)

[19-Remove-Nth-Node-From-End-of-List.md](Two-Pointer/19-Remove-Nth-Node-From-End-of-List/19-Remove-Nth-Node-From-End-of-List.md)

[876-Middle-of-the-Linked-List.md](Two-Pointer/876-Middle-of-the-Linked-List/876-Middle-of-the-Linked-List.md)

[142-Linked-List-Cycle-II.md](Two-Pointer/142-Linked-List-Cycle-II/142-Linked-List-Cycle-II.md)

[160-Intersection-of-Two-Linked-Lists.md](Two-Pointer/160-Intersection-of-Two-Linked-Lists/160-Intersection-of-Two-Linked-Lists.md)

[26-Remove-Duplicates-from-Sorted-Array.md](Two-Pointer/26-Remove-Duplicates-from-Sorted-Array/26-Remove-Duplicates-from-Sorted-Array.md)

[83-Remove-Duplicates-from-Sorted-List.md](Two-Pointer/83-Remove-Duplicates-from-Sorted-List/83-Remove-Duplicates-from-Sorted-List.md)

[283-Move-Zeroes.md](Two-Pointer/283-Move-Zeroes/283-Move-Zeroes.md)

[5-Longest-Palindromic-Substring.md](Two-Pointer/5-Longest-Palindromic-Substring/5-Longest-Palindromic-Substring.md)

[141-Linked-List-Cycle.md](Two-Pointer/141-Linked-List-Cycle/141-Linked-List-Cycle.md)

[82-Remove-Duplicates-from-Sorted-List-II.md](Two-Pointer/82-Remove-Duplicates-from-Sorted-List-II/82-Remove-Duplicates-from-Sorted-List-II.md)

[1836-Remove-Duplicates-from-Unsorted-List.md](Two-Pointer/1836-Remove-Duplicates-from-Unsorted-List/1836-Remove-Duplicates-from-Unsorted-List.md)

[378-Kth-Smallest-Element-in-a-Sorted-Matrix.md](Two-Pointer/378-Kth-Smallest-Element-in-a-Sorted-Matrix/378-Kth-Smallest-Element-in-a-Sorted-Matrix.md)

[373-Find-K-Pairs-with-Smallest-Sums.md](Two-Pointer/373-Find-K-Pairs-with-Smallest-Sums/373-Find-K-Pairs-with-Smallest-Sums.md)

[234-Palindrome-Linked-List.md](Two-Pointer/234-Palindrome-Linked-List/234-Palindrome-Linked-List.md)

[977-Squares-of-a-Sorted-Array.md](Two-Pointer/977-Squares-of-a-Sorted-Array/977-Squares-of-a-Sorted-Array.md)

[25-Reverse-Nodes-in-k-Group.md](Code/25-Reverse-Nodes-in-k-Group/25-Reverse-Nodes-in-k-Group.md)

[15-3Sum.md](Two-Pointer/15-3Sum/15-3Sum.md)

[18-4Sum.md](Two-Pointer/18-4Sum/18-4Sum.md)

[303-Range-Sum-Query-Immutable.md](Code/303-Range-Sum-Query-Immutable/303-Range-Sum-Query-Immutable.md)

[304-Range-Sum-Query-2D-Immutable.md](Code/304-Range-Sum-Query-2D-Immutable/304-Range-Sum-Query-2D-Immutable.md)

[1314-Matrix-Block-Sum.md](Code/1314-Matrix-Block-Sum/1314-Matrix-Block-Sum.md)

[525-Contiguous-Array.md](Code/525-Contiguous-Array/525-Contiguous-Array.md)

[523-Continuous-Subarray-Sum.md](Code/523-Continuous-Subarray-Sum/523-Continuous-Subarray-Sum.md)

[560-Subarray-Sum-Equals-K.md](Code/560-Subarray-Sum-Equals-K/560-Subarray-Sum-Equals-K.md)

[325-maximum-size-subarray-sum-equals-k.md](Code/325-maximum-size-subarray-sum-equals-k/325-maximum-size-subarray-sum-equals-k.md)

[974-Subarray-Sums-Divisible-by-K.md](Code/974-Subarray-Sums-Divisible-by-K/974-Subarray-Sums-Divisible-by-K.md)

[1124-Longest-Well-Performing-Interval.md](Code/1124-Longest-Well-Performing-Interval/1124-Longest-Well-Performing-Interval.md)

[370-Range-Addition.md](Code/370-Range-Addition/370-Range-Addition.md)

[1109-Corporate-Flight-Bookings.md](Code/1109-Corporate-Flight-Bookings/1109-Corporate-Flight-Bookings.md)

[76-Minimum-Window-Substring.md](Two-Pointer/76-Minimum-Window-Substring/76-Minimum-Window-Substring.md)

[567-Permutation-in-String.md](Two-Pointer/567-Permutation-in-String/567-Permutation-in-String.md)

[438-Find-All-Anagrams-in-a-String.md](Two-Pointer/438-Find-All-Anagrams-in-a-String/438-Find-All-Anagrams-in-a-String.md)

[3-Longest-Substring-Without-Repeating-Characters.md](Two-Pointer/3-Longest-Substring-Without-Repeating-Characters/3-Longest-Substring-Without-Repeating-Characters.md)

[1658-Minimum-Operations-to-Reduce-X-to-Zero.md](Two-Pointer/1658-Minimum-Operations-to-Reduce-X-to-Zero/1658-Minimum-Operations-to-Reduce-X-to-Zero.md)

[713-Subarray-Product-Less-Than-K.md](Two-Pointer/713-Subarray-Product-Less-Than-K/713-Subarray-Product-Less-Than-K.md)

[1004-Max-Consecutive-Ones-III.md](Two-Pointer/1004-Max-Consecutive-Ones-III/1004-Max-Consecutive-Ones-III.md)

[219-Contains-Duplicate-II.md](Two-Pointer/219-Contains-Duplicate-II/219-Contains-Duplicate-II.md)

[187-Repeated-DNA-Sequences.md](Two-Pointer/187-Repeated-DNA-Sequences/187-Repeated-DNA-Sequences.md)

[875-Koko-Eating-Bananas.md](Code/875-Koko-Eating-Bananas/875-Koko-Eating-Bananas.md)

[1011-Capacity-To-Ship-Packages-Within-D-Days.md](Code/1011-Capacity-To-Ship-Packages-Within-D-Days/1011-Capacity-To-Ship-Packages-Within-D-Days.md)

[410-Split-Array-Largest-Sum.md](Code/410-Split-Array-Largest-Sum/410-Split-Array-Largest-Sum.md)

[74-Search-a-2D-Matrix.md](Code/74-Search-a-2D-Matrix/74-Search-a-2D-Matrix.md)

[240-Search-a-2D-Matrix-II.md](Code/240-Search-a-2D-Matrix-II/240-Search-a-2D-Matrix-II.md)

[658-Find-K-Closest-Elements.md](Code/658-Find-K-Closest-Elements/658-Find-K-Closest-Elements.md)

[162-Find-Peak-Element.md](Code/162-Find-Peak-Element/162-Find-Peak-Element.md)

[852-Peak-Index-in-a-Mountain-Array.md](Code/852-Peak-Index-in-a-Mountain-Array/852-Peak-Index-in-a-Mountain-Array.md)

[LCR-173-takeAttendance.md](Code/LCR-173-takeAttendance/LCR-173-takeAttendance.md)

[33-Search-in-Rotated-Sorted-Array.md](Code/33-Search-in-Rotated-Sorted-Array/33-Search-in-Rotated-Sorted-Array.md)

[81-Search-in-Rotated-Sorted-Array-II.md](Code/81-Search-in-Rotated-Sorted-Array-II/81-Search-in-Rotated-Sorted-Array-II.md)

[528-Random-Pick-with-Weight.md](Code/528-Random-Pick-with-Weight/528-Random-Pick-with-Weight.md)

[870-Advantage-Shuffle.md](Code/870-Advantage-Shuffle/870-Advantage-Shuffle.md)

[380-Insert-Delete-GetRandom-O1.md](Code/380-Insert-Delete-GetRandom-O1/380-Insert-Delete-GetRandom-O1.md)

[710-Random-Pick-with-Blacklist.md](Code/710-Random-Pick-with-Blacklist/710-Random-Pick-with-Blacklist.md)

[1081-Smallest-Subsequence-of-Distinct-Characters.md](Code/1081-Smallest-Subsequence-of-Distinct-Characters/1081-Smallest-Subsequence-of-Distinct-Characters.md)

[316-Remove-Duplicate-Letters.md](Code/316-Remove-Duplicate-Letters/316-Remove-Duplicate-Letters.md)

---

## Other Skills

[204-Count-Primes.md](Code/204-Count-Primes/204-Count-Primes.md)

[264-Ugly-Number-II.md](Two-Pointer/264-Ugly-Number-II/264-Ugly-Number-II.md)

[48-Rotate-Image.md](Code/48-Rotate-Image/48-Rotate-Image.md)

[54-Spiral-Matrix.md](Code/54-Spiral-Matrix/54-Spiral-Matrix.md)

[59-Spiral-Matrix-II.md](Code/59-Spiral-Matrix-II/59-Spiral-Matrix-II.md)

---

## leetcode Practices and Memos:

1. [1-Two-Sum.md](Memos/1-Two-Sum.md)
2. [2-Add-Two-Number.md](Memos/2-Add-Two-Number.md)
3. [1480-Running-Sum-of-1d-Array.md](Memos/1480-Running-Sum-of-1d-Array.md)
4. [1672-Richest-Customer-Wealth.md](Memos/1672-Richest-Customer-Wealth.md)
5. [412-Fizz-Buzz.md](Memos/412-Fizz-Buzz.md)
6. [876-Middle-of-the-Linked-List.md](Array-and-LinkedList/876-Middle-of-the-Linked-List/876-Middle-of-the-Linked-List.md)
7. [383-Ransom-Note.md](Memos/383-Ransom-Note.md)
8. [88-Merge-Sorted-Array.md](Memos/88-Merge-Sorted-Array.md)
9. [27-Remove-Element.md](Memos/27-Remove-Element.md)
10. [26-Remove-Duplicates-from-Sorted-Array.md](Memos/26-Remove-Duplicates-from-Sorted-Array.md)
11. [80-Remove-Duplicates-from-Sorted-Array-II.md](Memos/80-Remove-Duplicates-from-Sorted-Array-II.md)
12. [169-Majority-Element.md](Memos/169-Majority-Element.md)
13. [189-Rotate-Array.md](Memos/189-Rotate-Array.md)
14. [121-Best-Time-to-Buy-and-Sell-Stock.md](Memos/121-Best-Time-to-Buy-and-Sell-Stock.md)
15. [122-Best-Time-to-Buy-and-Sell-Stock-II.md](Memos/122-Best-Time-to-Buy-and-Sell-Stock-II.md)
16. [55-Jump-Game.md](Memos/55-Jump-Game.md)
17. [45-Jump-Game-II.md](Memos/45-Jump-Game-II.md)
18. [1768-Merge-Strings-Alternately.md](Memos/1768-Merge-Strings-Alternately.md)
19. [1071-Greatest-Common-Divisor-of-Strings.md](Memos/1071-Greatest-Common-Divisor-of-Strings.md)
20. [1431-Kids-With-the-Greatest-Number-of-Candies.md](Memos/1431-Kids-With-the-Greatest-Number-of-Candies.md)
21. [605-Can-Place-Flowers.md](Memos/605-Can-Place-Flowers.md)
22. [345-Reverse-Vowels-of-a-String.md](Memos/345-Reverse-Vowels-of-a-String.md)
23. [217-Contains-Duplicate.md](Memos/217-Contains-Duplicate.md)
24. [53-Maximum-Subarray.md](Memos/53-Maximum-Subarray.md)
25. [Longest-Word.md](Memos/Longest-Word.md)
26. [283-Move-Zeroes.md](Two-Pointer/283-Move-Zeroes/283-Move-Zeroes.md)
27. [\* 1915-Number-of-Wonderful-Substrings.md](Memos/1915-Number-of-Wonderful-Substrings.md)
28. [2215-Find-the-Difference-of-Two-Arrays.md](Memos/2215-Find-the-Difference-of-Two-Arrays.md)
29. [1207-Unique-Number-of-Occurrences.md](Memos/1207-Unique-Number-of-Occurrences.md)
30. [392-Is-Subsequence.md](Memos/392-Is-Subsequence.md)
31. [Merge-Sorted-Arrays.md](Memos/Merge-Sorted-Arrays.md)
32. [151-Reverse-Words-in-a-String.md](Memos/151-Reverse-Words-in-a-String.md)
33. [238-Product-of-Array-Except-Self.md](Memos/238-Product-of-Array-Except-Self.md)
34. [334-Increasing-Triplet-Subsequence.md](Memos/334-Increasing-Triplet-Subsequence.md)
35. [\* 443-String-Compression.md](Memos/443-String-Compression.md)
36. [2441-Largest-Positive-Integer-That-Exists-With-Its-Negative.md](Memos/2441-Largest-Positive-Integer-That-Exists-With-Its-Negative.md)
37. [\* 322-Coin-Change.md](Memos/322-Coin-Change.md)
38. [881-Boats-to-Save-People.md](Memos/881-Boats-to-Save-People.md)
39. [11-Container-With-Most-Water.md](Memos/11-Container-With-Most-Water.md)
40. [1679-Max-Number-of-K-Sum-Pairs.md](Memos/1679-Max-Number-of-K-Sum-Pairs.md)
41. [237-Delete-Node-in-a-Linked-List.md](Memos/237-Delete-Node-in-a-Linked-List.md)
42. [328-Odd-Even-Linked-List.md](Memos/328-Odd-Even-Linked-List.md)
43. [206-Reverse-Linked-List.md](Memos/206-Reverse-Linked-List.md)
44. [2130-Maximum-Twin-Sum-of-a-Linked-List.md](Memos/2130-Maximum-Twin-Sum-of-a-Linked-List.md)
45. [2487-Remove-Nodes-From-Linked-List.md](Memos/2487-Remove-Nodes-From-Linked-List.md)
46. [2816-Double-a-Number-Represented-as-a-Linked-List.md](Memos/2816-Double-a-Number-Represented-as-a-Linked-List.md)
47. [1732-Find-the-Highest-Altitude.md](Memos/1732-Find-the-Highest-Altitude.md)
48. [724-Find-Pivot-Index.md](Memos/724-Find-Pivot-Index.md)
49. [506-Relative-Ranks.md](Memos/506-Relative-Ranks.md)
50. [21-Merge-Two-Sorted-Lists.md](Memos/21-Merge-Two-Sorted-Lists.md)
51. [389-Find-the-Difference.md](Memos/389-Find-the-Difference.md)
52. [28-Find-the-Index-of-the-First-Occurrence-in-a-String.md](Code/28-Find-the-Index-of-the-First-Occurrence-in-a-String/28-Find-the-Index-of-the-First-Occurrence-in-a-String.md)
53. [242-Valid-Anagram.md](Memos/242-Valid-Anagram.md)
54. [459-Repeated-Substring-Pattern.md](Memos/459-Repeated-Substring-Pattern.md)
55. [\* 3075-Maximize-Happiness-of-Selected-Children.md](Memos/3075-Maximize-Happiness-of-Selected-Children.md)
56. [66-Plus-One.md](Memos/66-Plus-One.md)
57. [1822-Sign-of-the-Product-of-an-Array.md](Memos/1822-Sign-of-the-Product-of-an-Array.md)
58. [1502-Can-Make-Arithmetic-Progression-From-Sequence.md](Memos/1502-Can-Make-Arithmetic-Progression-From-Sequence.md)
59. [896-Monotonic-Array.md](Memos/896-Monotonic-Array.md)
60. [13-Roman-to-Integer.md](Memos/13-Roman-to-Integer.md)
61. [58-Length-of-Last-Word.md](Memos/58-Length-of-Last-Word.md)
62. [709-To-Lower-Case.md](Memos/709-To-Lower-Case.md)
63. [445-Add-Two-Numbers-II.md](Memos/445-Add-Two-Numbers-II.md)
64. [682-Baseball-Game.md](Memos/682-Baseball-Game.md)
65. [657-Robot-Return-to-Origin.md](Memos/657-Robot-Return-to-Origin.md)
66. [1275-Find-Winner-on-a-Tic-Tac-Toe-Game.md](Memos/1275-Find-Winner-on-a-Tic-Tac-Toe-Game.md)
67. [1041-Robot-Bounded-In-Circle.md](Memos/1041-Robot-Bounded-In-Circle.md)
68. [643-Maximum-Average-Subarray-I.md](Memos/643-Maximum-Average-Subarray-I.md)
69. [2373-Largest-Local-Values-in-a-Matrix.md](Memos/2373-Largest-Local-Values-in-a-Matrix.md)
70. [861-Score-After-Flipping-Matrix.md](Code/861-Score-After-Flipping-Matrix/861-Score-After-Flipping-Matrix.md)
71. [700-Search-in-a-Binary-Search-Tree.md](Code/700-Search-in-a-Binary-Search-Tree/700-Search-in-a-Binary-Search-Tree.md)
72. [450-Delete-Node-in-a-BST.md](Code/450-Delete-Node-in-a-BST/450-Delete-Node-in-a-BST.md)
73. [2390-Removing-Stars-From-a-String.md](Code/2390-Removing-Stars-From-a-String/2390-Removing-Stars-From-a-String.md)
74. [1657-Determine-if-Two-Strings-Are-Close.md](Code/1657-Determine-if-Two-Strings-Are-Close/1657-Determine-if-Two-Strings-Are-Close.md)
75. [735-Asteroid-Collision.md](Code/735-Asteroid-Collision/735-Asteroid-Collision.md)
76. [394-Decode-String.md](Code/394-Decode-String/394-Decode-String.md)
77. [933-Number-of-Recent-Calls.md](Code/933-Number-of-Recent-Calls/933-Number-of-Recent-Calls.md)
78. [1325-Delete-Leaves-With-a-Given-Value.md](Code/1325-Delete-Leaves-With-a-Given-Value/1325-Delete-Leaves-With-a-Given-Value.md)
79. [649-Dota2-Senate.md](Code/649-Dota2-Senate/649-Dota2-Senate.md)
80. [232-Implement-Queue-using-Stacks.md](Code/232-Implement-Queue-using-Stacks/232-Implement-Queue-using-Stacks.md)
81. [1456-Maximum-Number-of-Vowels-in-a-Substring-of-Given-Length.md](Code/1456-Maximum-Number-of-Vowels-in-a-Substring-of-Given-Length/1456-Maximum-Number-of-Vowels-in-a-Substring-of-Given-Length.md)
82. [78-Subsets.md](Code/78-Subsets/78-Subsets.md)
83. [1493-Longest-Subarray-of-1's-After-Deleting-One-Element.md](Code/1493-Longest-Subarray-of-1s-After-Deleting-One-Element/1493-Longest-Subarray-of-1's-After-Deleting-One-Element.md)
84. [125-Valid-Palindrome.md](Code/125-Valid-Palindrome/125-Valid-Palindrome.md)
85. [1523-Count-Odd-Numbers-in-an-Interval-Range.md](Code/1523-Count-Odd-Numbers-in-an-Interval-Range/1523-Count-Odd-Numbers-in-an-Interval-Range.md)
86. [1491-Average-Salary-Excluding-the-Minimum-and-Maximum-Salary.md](Code/1491-Average-Salary-Excluding-the-Minimum-and-Maximum-Salary/1491-Average-Salary-Excluding-the-Minimum-and-Maximum-Salary.md)
87. [34-Find-First-and-Last-Position-of-Element-in-Sorted-Array.md](Two-Pointer/34-Find-First-and-Last-Position-of-Element-in-Sorted-Array/34-Find-First-and-Last-Position-of-Element-in-Sorted-Array.md)
88. [704-Binary-Search.md](Code/704-Binary-Search/704-Binary-Search.md)
89. [274-H-Index.md](Code/274-H-Index/274-H-Index.md)
90. [205-Isomorphic-Strings.md](Code/205-Isomorphic-Strings/205-Isomorphic-Strings.md)
91. [225-Implement-Stack-using-Queues.md](Code/225-Implement-Stack-using-Queues/225-Implement-Stack-using-Queues.md)
92. [143-Reorder-List.md](Code/143-Reorder-List/143-Reorder-List.md)
93. [20-Valid-Parentheses.md](Code/20-Valid-Parentheses/20-Valid-Parentheses.md)
94. [150-Evaluate-Reverse-Polish-Notation.md](Code/150-Evaluate-Reverse-Polish-Notation/150-Evaluate-Reverse-Polish-Notation.md)
95. [388-Longest-Absolute-File-Path.md](Code/388-Longest-Absolute-File-Path/388-Longest-Absolute-File-Path.md)
96. [155-Min-Stack.md](Code/155-Min-Stack/155-Min-Stack.md)
97. [239-Sliding-Window-Maximum.md](Code/239-Sliding-Window-Maximum/239-Sliding-Window-Maximum.md)
98. [622-Design-Circular-Queue.md](Code/622-Design-Circular-Queue/622-Design-Circular-Queue.md)
99. [641-Design-Circular-Deque.md](Code/641-Design-Circular-Deque/641-Design-Circular-Deque.md)
100. [1670-Design-Front-Middle-Back-Queue.md](Code/1670-Design-Front-Middle-Back-Queue/1670-Design-Front-Middle-Back-Queue.md)
101. [496-Next-Greater-Element-I.md](Code/496-Next-Greater-Element-I/496-Next-Greater-Element-I.md)
102. [503-Next-Greater-Element-II.md](Code/503-Next-Greater-Element-II/503-Next-Greater-Element-II.md)
103. [556-Next-Greater-Element-III.md](Code/556-Next-Greater-Element-III/556-Next-Greater-Element-III.md)
104. [1019-Next-Greater-Node-In-Linked-List.md](Code/1019-Next-Greater-Node-In-Linked-List/1019-Next-Greater-Node-In-Linked-List.md)
105. [1944-Number-of-Visible-People-in-a-Queue.md](Code/1944-Number-of-Visible-People-in-a-Queue/1944-Number-of-Visible-People-in-a-Queue.md)
106. [1475-Final-Prices-With-a-Special-Discount-in-a-Shop.md](Code/1475-Final-Prices-With-a-Special-Discount-in-a-Shop/1475-Final-Prices-With-a-Special-Discount-in-a-Shop.md)
107. [901-Online-Stock-Span.md](Code/901-Online-Stock-Span/901-Online-Stock-Span.md)
108. [402-Remove-K-Digits.md](Code/402-Remove-K-Digits/402-Remove-K-Digits.md)
109. [1438-Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit.md](Code/1438-Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit/1438-Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit.md)
110. [862-Shortest-Subarray-with-Sum-at-Least-K.md](Code/862-Shortest-Subarray-with-Sum-at-Least-K/862-Shortest-Subarray-with-Sum-at-Least-K.md)
111. [135-Candy.md](Code/135-Candy/135-Candy.md)
112. [918-Maximum-Sum-Circular-Subarray.md](Code/918-Maximum-Sum-Circular-Subarray/918-Maximum-Sum-Circular-Subarray.md)
113. [1696-Jump-Game-VI.md](Code/1696-Jump-Game-VI/1696-Jump-Game-VI.md)
114. [1425-Constrained-Subsequence-Sum.md](Code/1425-Constrained-Subsequence-Sum/1425-Constrained-Subsequence-Sum.md)
115. [146-LRU-Cache.md](Code/146-LRU-Cache/146-LRU-Cache.md)
116. [134-Gas-Station.md](Code/134-Gas-Station/134-Gas-Station.md)
117. [209-Minimum-Size-Subarray-Sum.md](Code/209-Minimum-Size-Subarray-Sum/209-Minimum-Size-Subarray-Sum.md)
118. [290-Word-Pattern.md](Code/290-Word-Pattern/290-Word-Pattern.md)
119. [49-Group-Anagrams.md](Code/49-Group-Anagrams/49-Group-Anagrams.md)
120. [12-Integer-to-Roman.md](Code/12-Integer-to-Roman/12-Integer-to-Roman.md)
121. [6-Zigzag-Conversion.md](Code/6-Zigzag-Conversion/6-Zigzag-Conversion.md)
122. [7-Reverse-Integer.md](Code/7-Reverse-Integer/7-Reverse-Integer.md)
123. [8-String-to-Integer-(atoi).md)](<Code/8-String-to-Integer-atoi/8-String-to-Integer-(atoi).md>)
124. [9-Palindrome-Number.md](Code/9-Palindrome-Number/9-Palindrome-Number.md)
125. [29-Divide-Two-Integers.md](Code/29-Divide-Two-Integers/29-Divide-Two-Integers.md)
126. [202-Happy-Number.md](Code/202-Happy-Number/202-Happy-Number.md)
127. [2053-Kth-Distinct-String-in-an-Array.md](Code/2053-Kth-Distinct-String-in-an-Array/2053-Kth-Distinct-String-in-an-Array.md)
128. [228-Summary-Ranges.md](Code/228-Summary-Ranges/228-Summary-Ranges.md)
129. [2300-Successful-Pairs-of-Spells-and-Potions.md](Binary-Tree/2300-Successful-Pairs-of-Spells-and-Potions/2300-Successful-Pairs-of-Spells-and-Potions.md)
130. [35-Search-Insert-Position.md](Code/35-Search-Insert-Position/35-Search-Insert-Position.md)
131. [67-Add-Binary.md](Code/67-Add-Binary/67-Add-Binary.md)
132. [128-Longest-Consecutive-Sequence.md](Code/128-Longest-Consecutive-Sequence/128-Longest-Consecutive-Sequence.md)
133. [190-Reverse-Bits.md](Code/190-Reverse-Bits/190-Reverse-Bits.md)
134. [860-Lemonade-Change.md](Code/860-Lemonade-Change/860-Lemonade-Change.md)
135. [118-Pascal's-Triangle.md](Code/118-Pascal's-Triangle/118-Pascal's-Triangle.md)

---

## Paiza:

1. [C084.md](Paiza/C/C084/C084.md)
2. [B138.md](Paiza/B138/B138.md)
3. [Search-History.md](Paiza/C/Search-History/Search-History.md)
4. [lottery.md](Paiza/C/lottery/lottery.md)
5. [Word-Count-Part1.md](Paiza/C/Word-Count-Part1/Word-Count-Part1.md)
6. [Word-Count-Part2.md](Paiza/C/Word-Count-Part2/Word-Count-Part2.md)
7. [Word-Count-Part3.md](Paiza/C/Word-Count-Part3/Word-Count-Part3.md)
8. [Word-Count-Part4.md](Paiza/C/Word-Count-Part4/Word-Count-Part4.md)
9. [Word-Count-Part5.md](Paiza/C/Word-Count-Part5/Word-Count-Part5.md)
10. [Word-Count-Part6.md](Paiza/C/Word-Count-Part6/Word-Count-Part6.md)
11. [Word-Count-Part7.md](Paiza/C/Word-Count-Part7/Word-Count-Part7.md)
12. [Word-Count-final.md](Paiza/C/Word-Count-final/Word-Count-final.md)
13. [Orange-Divide.md](Paiza/C/Orange-Divide/Orange-Divide.md)
14. [Sort.md](Paiza/C/Sort/Sort.md)
15. [C097.md](Paiza/C/C097/C097.md)
16. [C139.md](Paiza/C/C139/C139.md)
17. [C120.md](Paiza/C/C120/C120.md)
18. [C030.md](Paiza/C/C030/C030.md)
19. [C043.md](Paiza/C/C043/C043.md)
20. [C090.md](Paiza/C/C090/C090.md)
21. [B029.md](Paiza/B/B029/B029.md)
22. [B131.md](Paiza/B/B131/B131.md)
23. [A066.md](Paiza/A/A066/A066.md)
24. [A081.md](Paiza/A/A081/A081.md)

---
